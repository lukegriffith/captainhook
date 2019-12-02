package captainhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Structure is responsible to processing incoming requests against configured endpoints.
type HookEngine struct {
	log         *log.Logger
	endpointSvc EndpointService
	secretEng   SecretEngine
}

// Constructor requires endpoint and secret engine.
func NewHookEngine(log *log.Logger, ec EndpointService, sec SecretEngine) *HookEngine {
	return &HookEngine{log, ec, sec}
}

// Main routine that processes recieved hooks, obtaining endpoints and processing rules.
// Various error checking and validation happens at this stage, i.e mapping required secrets to
// dataBag. Databag is a map of input parameters passed to each rules function.
func (h *HookEngine) Hook(w http.ResponseWriter, r *http.Request) {

	var name string
	var ok bool

	// Extract variables from request.
	vars := mux.Vars(r)

	if name, ok = vars["id"]; !ok {
		w.WriteHeader(http.StatusNotFound)
		h.log.Println("no id provided.")
		return
	}
	h.log.Println("processing webhook:", name)

	var endpoint *Endpoint



	endpoint, err := h.endpointSvc.Endpoint(name)


	if err != nil {
		h.log.Println("error getting endpoint", name)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}



	b, err := ioutil.ReadAll(r.Body)



	if err != nil {
		h.log.Println("unable to get body from request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body := fmt.Sprintf("%s", b)
	decodedBody, err := url.QueryUnescape(body)

	if err != nil {
		h.log.Fatal("unable to URL decode body")
	}

	dataBag := make(map[string]interface{})
	err = json.Unmarshal([]byte(decodedBody), &dataBag)

	if err != nil {
		h.log.Println("unable to unmarshal json")
	}

	var secretMap map[string]string = make(map[string]string)

	for _, secret := range endpoint.Secrets {
		v, err := h.secretEng.GetTextSecret(secret)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			h.log.Println("unable to get secret from engine", secret)
			return
		}

		secretMap[secret] = v
	}

	dataBag["_secrets"] = secretMap


	h.executeEndpoint(endpoint, r, w, &dataBag)


}

// Executes all endpoint rules, returns bool to determine if
func (h *HookEngine) executeEndpoint(e *Endpoint, r *http.Request, w http.ResponseWriter,  dataBag *map[string]interface{})  {

	var request bytes.Buffer
	var echoStrings []string

	rules, err := e.GetRules()

	if err != nil {
		h.log.Println("unable to enumerate rules, endpoint", e.Name)
		w.WriteHeader(http.StatusInternalServerError)
	}

	for _, r := range rules {

		AssignFunction(&r)

		err = r.Execute(&request, *dataBag)

		if err != nil {
			h.log.Println(r, "failed to execute template.", err)
			continue
		}
		h.log.Println("rendered template: ", request.String())


		if r.Destination != "" {

			h.log.Println("forwarding to", r.Destination)
			_, err = http.Post(r.Destination, "application/json", &request)

			if err != nil {
				h.log.Println("post request to", r.Destination, "failed.")
			}
		}

		if r.Echo {
			echoStrings = append(echoStrings, request.String())
		}
		request.Reset()
	}

	if len(echoStrings) > 0 {
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(strings.Join(echoStrings[:], "\n")))

		if err != nil {
			h.log.Println("Unable to echo reply.")
		}
	} else {
		w.WriteHeader(http.StatusNoContent)
	}


}
