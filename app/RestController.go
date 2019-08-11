package app

import (
	"log"
	"net/http"
)

type RestController interface{
	Post(w http.ResponseWriter, r *http.Request, l *log.Logger)
	Get(w http.ResponseWriter, r *http.Request, l *log.Logger)
	Patch(w http.ResponseWriter, r *http.Request, l *log.Logger)
	Delete(w http.ResponseWriter, r *http.Request, l *log.Logger)
}


