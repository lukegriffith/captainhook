package server

import (
	"github.com/gorilla/mux"
	"github.com/lukemgriffith/captainhook"
	"github.com/lukemgriffith/captainhook/util"
	"log"
	"net/http"
)

//TODO: Document
func New(endpoints captainhook.EndpointService, secrets captainhook.SecretEngine) http.Handler {

	log := util.NewLog("CaptainHook ")
	mux := mux.NewRouter()
	log.Println("Starting Application Server.")
	AppServer := &AppServer{mux, log}

	webhookController := NewRestController(NewWebhookController(endpoints, secrets))
	mux.HandleFunc("/webhook/{id}", webhookController.ServeHTTP)

	endpointController := NewRestController(NewEndpointController(endpoints))
	mux.HandleFunc("/endpoint", endpointController.ServeHTTP)
	mux.HandleFunc("/endpoint/{name}", endpointController.ServeHTTP)

	return AppServer
}

//TODO: Document
type AppServer struct {
	mux *mux.Router
	log *log.Logger
}

//TODO: Document
func (a *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
