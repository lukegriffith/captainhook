package captainhook

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type HookEngine struct {
	secret string
	log *log.Logger
	endpointSvc EndpointService
}

func NewHookEngine(secret string, log *log.Logger, ec *EndpointService) *HookEngine{
	return &HookEngine{secret, log, *ec}
}

// TODO
//  This is likely blocking the main execution thread. I think i'll need to send details to a channel and have a
//  separate go routine pickup the request.

func (h *HookEngine) Hook(w http.ResponseWriter, r *http.Request) {

	var name string
	var ok bool

	vars := mux.Vars(r)

	if name, ok = vars["id"]; ! ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	endpoints, err :=  h.endpointSvc.Endpoints()

	if err != nil {
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

	if ! found {
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

	for _, r := range rules {
		r.Execute(decodedBody)
	}

	w.WriteHeader(http.StatusNoContent)
}
