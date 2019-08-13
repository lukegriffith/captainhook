package server

import (
  "github.com/lukemgriffith/captainhook"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
	"fmt"
	"log"
)

func New(es *captainhook.EndpointService) http.Handler {

	log := NewLog("CaptainHook")
	mux := mux.NewRouter()
	fs := http.FileServer(http.Dir("static"))
	AppServer := &AppServer{mux, log}

	log.Println("Starting AppServerlication.")

	mux.Handle("/", fs)
	mux.HandleFunc("/webhook/{id}", AppServer.hook)

	ec := NewRestController(NewEndpointController(es))

	mux.HandleFunc("/endpoint/", ec.ServeHTTP)
	mux.HandleFunc("/endpoint/{name}", ec.ServeHTTP)

	return AppServer

}

type AppServer struct {
	mux *mux.Router
	log *log.Logger
}

func (a *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}

func (a *AppServer) hook(w http.ResponseWriter, r *http.Request) {

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
