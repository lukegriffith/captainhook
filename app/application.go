package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func New() http.Handler {

	log := log.New(os.Stdout, "app", log.LstdFlags)
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("static"))
	app := &app{mux, log}

	log.Println("Starting application.")

	mux.Handle("/", fs)
	mux.HandleFunc("/webhook/", app.hooks)

	ec := NewEndpointController()

	mux.HandleFunc("/endpoint/", ec.Serve)

	return app

}

type app struct {
	mux *http.ServeMux
	log *log.Logger
}

func (a *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}

func (a *app) hooks(w http.ResponseWriter, r *http.Request) {

	var secret string

	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		a.log.Fatal("Unable to get body from request")
	}

	body := fmt.Sprintf("%s", b)

	secret = r.Header.Get("secret")

	decodedBody, err := url.QueryUnescape(body)

	if err != nil {
		a.log.Fatal("Unable to URL decode body")
	}

	a.log.Print(decodedBody, " ", secret, " ", r.URL)

	w.WriteHeader(204)
}

func NewLog(name string) *log.Logger {
	return log.New(os.Stdout, name, log.LstdFlags)
}
