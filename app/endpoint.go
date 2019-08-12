package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Endpoint struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Secret string   `json:"secret"`
	Rules  []Rule   `json:"rules"`
	Source []Source `json:"sources"`
}

type EndpointController struct {
	log *log.Logger
}

func NewEndpointController() *EndpointController {
	log := NewLog("EndpointController")
	return &EndpointController{log}
}

// Get recieved a single instance of Endpoint
func (e *EndpointController) Get(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	e.log.Println(vars)

	end := Endpoint{"1", "test", "testsec", nil, nil}
	json, err := json.Marshal(end)
	if err != nil {
		e.log.Fatal("Unable to convert endpoint to json")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func (e *EndpointController) Post(w http.ResponseWriter, r *http.Request)   {}
func (e *EndpointController) Patch(w http.ResponseWriter, r *http.Request)  {}
func (e *EndpointController) Delete(w http.ResponseWriter, r *http.Request) {}
