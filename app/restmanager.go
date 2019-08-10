package app

import (
	"log"
	"net/http"
)

type restManager struct {
	handler handler
	log     *log.Logger
}

func (rm *restManager) serve(w http.ResponseWriter, r *http.Request) {
	rm.log.Println("Rest Manager recieved call.")

	if r.Method == "GET" {
		rm.handler.Get(w, r, rm.log)
	} else {
		w.WriteHeader(405)
	}
	/* else if r.Method == "POST" {
	  rm.handler.Post(w, r, rm.log)
	} else if r.Method == "PATCH" {
	  rm.handler.Update(w, r, rm.log)
	}*/
}
