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
	"text/template"
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

// Main routine that processes received hooks, obtaining endpoints and processing rules.
// Various error checking and validation happens at this stage, i.e mapping required secrets to
// dataBag. data bag is a map of input parameters passed to each rules function.
func (h *HookEngine) Hook(w http.ResponseWriter, r *http.Request) {

	var dataBag = make(map[string]interface{})
	var secretMap = make(map[string]string)
	var endpoint *Endpoint
	var name string
	var ok bool

	// Extract variables from request.
	vars := mux.Vars(r)

	// Validate ID is provided.
	if name, ok = vars["id"]; !ok {
		w.WriteHeader(http.StatusNotFound)
		h.log.Println("no id provided.")
		return
	}
	h.log.Println("processing webhook:", name)

	// Get endpoint by identifier.
	endpoint, err := h.endpointSvc.Endpoint(name)

	if err != nil {
		h.log.Println("error getting endpoint", name)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Read body and create data bag string map
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

	err = json.Unmarshal([]byte(decodedBody), &dataBag)

	if err != nil {
		h.log.Println("unable to unmarshal json")
	}

	// Attach secrets to secrets map/
	for _, secret := range endpoint.Secrets {
		v, err := h.secretEng.GetTextSecret(secret)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			h.log.Println("unable to get secret from engine", secret)
			return
		}

		secretMap[secret] = v
	}

	// Store secrets on data bag.
	dataBag["_secrets"] = secretMap

	h.executeEndpoint(endpoint, r, w, &dataBag)
}

// Executes endpoint rules, returns data to caller on rules specified with echo.
func (h *HookEngine) executeEndpoint(e *Endpoint, r *http.Request, w http.ResponseWriter, dataBag *map[string]interface{}) {

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

			client := &http.Client{}

			dest := h.templateString(r.Destination, dataBag)

			h.log.Println("forwarding to", dest)

			req, err := http.NewRequest("POST", dest, &request)

			for k, v := range r.Headers {
				value := h.templateString(v, dataBag)
				req.Header.Add(k, value)
			}

			resp, err := client.Do(req)

			if err != nil {
				h.log.Println("post request to", r.Destination, "failed.")
			}

			h.log.Println(*dataBag)

			h.log.Println(resp)
		}

		if r.Echo {
			echoStrings = append(echoStrings, request.String())
		}
		request.Reset()
	}

	// Reply with rules that have been echoed
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

func (h *HookEngine) templateString(templ string, data *map[string]interface{}) string {

	tmpl, err := template.New("tmpl").Parse(templ)

	if err != nil {
		h.log.Println("Unable to create template for header from: ", templ)
		return ""
	}

	buf := make([]byte, 0, 1)
	var tpl *bytes.Buffer = bytes.NewBuffer(buf)
	err = tmpl.Execute(tpl, &data)
	return tpl.String()

}
