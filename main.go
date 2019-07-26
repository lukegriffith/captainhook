package main

import (
	"context"
	"github.com/lukemgriffith/captainhook/app"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	app := app.New()

	server := &http.Server{
		Addr:    ":8081",
		Handler: app,
	}

	go server.ListenAndServe()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	for {
		s := <-c
		log.Print("os signal recieved. processing.")

		switch s {
		case syscall.SIGTERM:
			log.Print("SIGTERM: shutting server down gracefully")
			server.Shutdown(context.Background())
			return
		}
	}
}
