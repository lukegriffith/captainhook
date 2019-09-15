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
	secret      string
	log         *log.Logger
	endpointSvc EndpointService
}

func NewHookEngine(secret string, log *log.Logger, ec *EndpointService) *HookEngine {
	return &HookEngine{secret, log, *ec}
}

// TODO
//  This is likely blocking the main execution thread. I think i'll need to send details to a channel and have a
//  separate go routine pickup the request.

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

	endpoints, err := h.endpointSvc.Endpoints()

	if err != nil {
    h.log.Println("Error getting endpoints.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var found bool = false
	var endpoint Endpoint

	for _, val := range endpoints {
		if val.Name == name {
			endpoint = val
			found = true
			break
		}
	}

	if !found {
    h.log.Println("Endpoint", name, "does not exist.")
		w.WriteHeader(http.StatusNotFound)
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
    h.log.Println(r)
    h.log.Println("Forwarding to", r.Destination_url)
    http.Post(r.Destination_url, "application/json", &request)
    request.Reset()
	}

	w.WriteHeader(http.StatusNoContent)
}
