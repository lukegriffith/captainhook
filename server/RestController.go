 package server

import (
	"github.com/lukemgriffith/captainhook/util"
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

func NewRestController(c Controller) RestController {
	log := util.NewLog("RestController ")
	return RestController{log, c}
}

//TODO: Document
func (rc *RestController) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if method := "GET"; r.Method == method {
		rc.log.Println(method, r.URL)
		rc.controller.Get(w, r)
	} else if method := "POST"; r.Method == method {
		rc.log.Println(method, r.URL)
		rc.controller.Post(w, r)
	} else if method := "PATCH"; r.Method == method {
		rc.log.Println(method, r.URL)
		rc.controller.Patch(w, r)
	} else if method := "DELETE"; r.Method == method {
		rc.log.Println(method, r.URL)
		rc.controller.Delete(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
