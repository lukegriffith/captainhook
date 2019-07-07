package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func handler(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal("Unable to get body from")
	}


	body := fmt.Sprintf("%s", b)

	decodedBody, err := url.QueryUnescape(body)

	if err != nil {
		log.Fatal("Unable to URL decode body")
	}

	log.Print(r)

	log.Print(decodedBody)

	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(200)
}

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))

}
