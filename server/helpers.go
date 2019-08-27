package server

import (
	"log"
	"os"
)

func NewLog(name string) *log.Logger {
	l := log.New(os.Stdout, name, log.LstdFlags)
	l.SetOutput(os.Stdout)
	return l
}
