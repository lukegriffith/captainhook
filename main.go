package main

import (
	"io/ioutil"
	"log"
	"net/http"
)


func handler(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal("Unable to get body from")
	}

	log.Printf("%s", b)

	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(200)
}

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))

}
