package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lukemgriffith/captainhook"
	"github.com/lukemgriffith/captainhook/util"
	"io/ioutil"
	"net/http"
	"net/url"
)

type WebhookController struct {
	endpoints captainhook.EndpointService
	secrets captainhook.SecretEngine
	log *util.Logger
}


func NewWebhookController(es captainhook.EndpointService, ss captainhook.SecretEngine) *WebhookController {
	log := util.NewLog("Webhook ")
	return &WebhookController{es, ss, log}
}

func (con *WebhookController) Post(w http.ResponseWriter, r *http.Request) {

	var dataBag = make(map[string]interface{})

	var endpoint *captainhook.Endpoint
	var name string
	var ok bool

	// Extract variables from request.
	vars := mux.Vars(r)

	// Validate ID is provided.
	if name, ok = vars["id"]; !ok {
		w.WriteHeader(http.StatusNotFound)
		con.log.Println("no id provided.")
		return
	}

	con.log.Println("processing webhook:", name)

	// Get endpoint by identifier.
	endpoint, err := con.endpoints.Endpoint(name)

	if err != nil {
		con.log.Println("error getting endpoint", name)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}


	// Read body and create data bag string map
	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		con.log.Println("unable to get body from request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body := fmt.Sprintf("%s", b)
	decodedBody, err := url.QueryUnescape(body)

	if err != nil {
		con.log.Fatal("unable to URL decode body")
	}

	err = json.Unmarshal([]byte(decodedBody), &dataBag)

	if err != nil {
		con.log.Println("unable to unmarshal json")
	}

	captainhook.Hook(w, r, endpoint, con.secrets, con.log, &dataBag)
}

func (con *WebhookController) Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (con *WebhookController) Patch(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (con *WebhookController) Delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}
