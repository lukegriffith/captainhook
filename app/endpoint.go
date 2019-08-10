package app

import (
	"encoding/json"
	"log"
	"net/http"
)

type endpoint struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Secret string `json:"secret"`
	Rules  []rule `json:"rules"`
	Source string `json:"source"`
}

func (e endpoint) Get(w http.ResponseWriter, r *http.Request, l *log.Logger) {
	json, err := json.Marshal(e)
	if err != nil {
		log.Fatal("Unable to convert endpoint to json")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
