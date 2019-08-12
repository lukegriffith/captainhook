package server

import (
  "log"
  "os"
)

func NewLog(name string) *log.Logger {
	return log.New(os.Stdout, name, log.LstdFlags)
}
