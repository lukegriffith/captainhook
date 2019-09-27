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
)

type HookEngine struct {
	log         *log.Logger
	endpointSvc EndpointService
}

func NewHookEngine(log *log.Logger, ec *EndpointService) *HookEngine {
	return &HookEngine{log, *ec}
}

func (h *HookEngine) Hook(w http.ResponseWriter, r *http.Request) {

	h.log.Println("processing webhook")

	var name string
	var ok bool

	vars := mux.Vars(r)

	if name, ok = vars["id"]; !ok {
		w.WriteHeader(http.StatusNotFound)

		h.log.Println("Unable to identify ID")
		return
	}

	var endpoint *Endpoint


	endpoint, err := h.endpointSvc.Endpoint(name)

	if err != nil {
		h.log.Println("Error getting endpoint", name)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		h.log.Println("Unable to get body from request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body := fmt.Sprintf("%s", b)
	decodedBody, err := url.QueryUnescape(body)

	//secret = r.Header.Get("secret")

	if err != nil {
		h.log.Fatal("Unable to URL decode body")
	}

	rules, err := endpoint.GetRules()

	if err != nil {
		h.log.Println("Unable to enumerate rules endpoint", endpoint.Name)
		w.WriteHeader(http.StatusInternalServerError)
	}

	dataBag := make(map[string]interface{})
	err = json.Unmarshal([]byte(decodedBody), &dataBag)

	if err != nil {
		h.log.Println("Unable to unmarshal json")
	}

	// TODO this area is not complete. request needs to be sent to destination
	// url.
	var request bytes.Buffer

	for _, r := range rules {
		r.Execute(&request, dataBag)
		h.log.Println("Rendered template: ", request.String())
		h.log.Println("Forwarding to", r.Destination)
		http.Post(r.Destination, "application/json", &request)
		request.Reset()
	}

	w.WriteHeader(http.StatusNoContent)
}
