package server

import (
	"log"
	"net/http"
)

type Controller interface {
	Post(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type RestController struct {
	log        *log.Logger
	controller Controller
}

func NewRestController(c Controller) RestController {
	log := NewLog("RestController ")
	return RestController{log, c}
}

func (rc *RestController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rc.log.Println("Rest Controller recieved call.")

	if r.Method == "GET" {
		rc.log.Println("Get method called on", r.URL)
		rc.controller.Get(w, r)
	} else if r.Method == "POST" {
		rc.log.Println("Post method called on", r.URL)
		rc.controller.Post(w, r)
	} else if r.Method == "PATCH" {
		rc.log.Println("Patch method called on", r.URL)
		rc.controller.Patch(w, r)
	} else if r.Method == "DELETE" {
		rc.log.Println("Delete method called on", r.URL)
		rc.controller.Delete(w, r)
	} else {
		w.WriteHeader(405)
	}

}