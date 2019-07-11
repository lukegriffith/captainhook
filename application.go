package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func New() http.Handler {

	mux := http.NewServeMux()
	log := log.New(os.Stdout, "web ", log.LstdFlags)
	app := &app{mux, log}

	mux.HandleFunc("/foo", app.foo)

	return app

}

type app struct {
	mux *http.ServeMux
	log *log.Logger
}

func (a *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}

func (a *app) foo(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		a.log.Fatal("Unable to get body from")
	}

	body := fmt.Sprintf("%s", b)

	decodedBody, err := url.QueryUnescape(body)

	if err != nil {
		a.log.Fatal("Unable to URL decode body")
	}

	parse(body)
	a.log.Print(r)

	a.log.Print(decodedBody)

	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(200)
}
