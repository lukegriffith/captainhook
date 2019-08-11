package app

import (
  "log"
  "net/http"
  "encoding/json"
)


type EndpointController struct {
  log *log.Logger
}

func NewEndpointController() *EndpointController {
  log := NewLog("endpoint controller")

  return &EndpointController{log}
}

func (c *EndpointController) Serve(w http.ResponseWriter, r *http.Request) {
	c.log.Println("Endpoint controller recieved call.")

	if r.Method == "GET" {
		c.Get(w, r)
	} else {
		w.WriteHeader(405)
	}
}

func (c *EndpointController) ServeAll(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		c.GetAll(w, r)
	} else {
		w.WriteHeader(405)
	}
}

// GetAll recieves all instances of Endpoint
func (e *EndpointController) GetAll(w http.ResponseWriter, r *http.Request) {

}


// Get recieved a single instance of Endpoint
func (e *EndpointController) Get(w http.ResponseWriter, r *http.Request) {
	end := endpoint{"1", "test", "testsec", nil, "sda"}
	json, err := json.Marshal(end)
	if err != nil {
		e.log.Fatal("Unable to convert endpoint to json")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
