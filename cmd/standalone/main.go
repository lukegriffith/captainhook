package main

import (
	"context"
	"fmt"
	"github.com/lukemgriffith/captainhook/configparser"
	"github.com/lukemgriffith/captainhook/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	configpath string = "config.yml"
	secretpath string = "secrets.yml"
)

func main() {

	svc, err := configparser.NewConfig(configpath)

	if err != nil {
		panic(err)
	}

	secEng, err := configparser.NewSecretEngine(secretpath, "test")

	if err != nil {
		panic(err)
	}

	endpoints, err := svc.Endpoints()

	if err != nil {
		panic(err)
	}


	// Validating all configured secrets are available.
	for _, end := range endpoints {
		for _, name := range end.Secrets {
			_, err := secEng.GetTextSecret(name)

			if err != nil {
				panic(fmt.Sprintf("unable to find secret %s", name))
			}
		}
	}


	app := server.New(svc, secEng)

	hookserver := &http.Server{
		Addr:    ":8081",
		Handler: app,
	}

	go hookserver.ListenAndServe()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	_ = <-c
	log.Print("os signal recieved processing.")

	log.Print("shutting server down gracefully.")
	hookserver.Shutdown(context.Background())

}
