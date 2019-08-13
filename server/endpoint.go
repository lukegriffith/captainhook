package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lukemgriffith/captainhook"
	"log"
	"net/http"
)

type EndpointController struct {
	service captainhook.EndpointService
	log     *log.Logger
}

func NewEndpointController(es captainhook.EndpointService) *EndpointController {
	log := NewLog("EndpointController ")
	return &EndpointController{es, log}
}

// Get recieved a single instance of Endpoint
func (e *EndpointController) Get(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	if name, ok := vars["name"]; ok {
		obj, err := e.service.Endpoint(name)

		if err != nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		json, err := json.Marshal(obj)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(json)
	} else {

		obj, err := e.service.Endpoints()

		if err != nil {
      w.WriteHeader(http.StatusInternalServerError)
      return
		}

		json, err := json.Marshal(obj)
    if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}


    w.WriteHeader(http.StatusOK)
    w.Write(json)
	}

}

func (e *EndpointController) Post(w http.ResponseWriter, r *http.Request)   {}
func (e *EndpointController) Patch(w http.ResponseWriter, r *http.Request)  {}
func (e *EndpointController) Delete(w http.ResponseWriter, r *http.Request) {}
