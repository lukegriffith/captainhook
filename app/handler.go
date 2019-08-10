package app

import (
  "net/http"
  "log"
)

type handler interface {
  //Post(w http.ResponseWriter, r *http.Request, l *log.Logger)
  Get(w http.ResponseWriter, r *http.Request, l *log.Logger)
  //Update(w http.ResponseWriter, r *http.Request, l *log.Logger)
}
