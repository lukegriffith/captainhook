package main

import (
	"context"
	"github.com/lukemgriffith/captainhook/configparser"
	"github.com/lukemgriffith/captainhook/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var configpath string = "config.yml"

func main() {

	_, svc := configparser.NewConfig(configpath)

	app := server.New(svc)

	server := &http.Server{
		Addr:    ":8081",
		Handler: app,
	}

	go server.ListenAndServe()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGUSR2)

	_ = <-c
	log.Print("os signal recieved processing.")

	log.Print("shutting server down gracefully.")
	server.Shutdown(context.Background())

}
