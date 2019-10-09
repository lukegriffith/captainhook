package server

import (
	"log"
	"net/http"
)

//TODO: Document
// Controller specifies the objectes required to participate in the server. an implementing object must comply to these
// methods to allow the server to correctly request its services.
type Controller interface {
	Post(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

//TODO: Holds a controller and logging structure, is mainly a way to get around go's lack of generics.
type RestController struct {
	log        *log.Logger
	controller Controller
}

//TODO: Document
func NewRestController(c Controller) RestController {
	log := NewLog("RestController ")
	return RestController{log, c}
}

//TODO: Document
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
