package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	alive := true

	app := New()

	server := &http.Server{
		Addr:         ":8081",
		Handler:      app,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go server.ListenAndServe()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	for alive {
		s := <-c
		alive = !interrupt(s)
		server.Shutdown(context.Background())
	}

}

func interrupt(sig os.Signal) bool {

	die := false

	switch sig {
	case syscall.SIGINT:
		log.Print("Interrupt recieved, starting graceful shutdown.")
		die = true

	default:
		log.Print("Unrecognized signal recieved. Ignoring")
	}

	return die
}
