package util

import (
	"log"
	"os"
)

var (
  debug bool = false
)

func SetDebug(enable bool) {
  debug = enable
}

type Logger struct {
  log *log.Logger
  verbose bool
}

func (l *Logger) Println(msg ...interface{}) {
  if debug {
    log.Println(msg)
  }
}

func (l *Logger) Print(msg ...interface{}) {
  if l.verbose || debug {
    l.log.Print(msg)
  }
}

func (l *Logger) Fatal(msg ...interface{}) {
  if l.verbose || debug {
    l.log.Fatal(msg)
  }
}

func NewDebugLog(name string) *Logger {
	l := log.New(os.Stdout, name, log.LstdFlags)
	l.SetOutput(os.Stdout)

	return &Logger{l, false}
}


//TODO: Document
func NewLog(name string) *Logger {
	l := log.New(os.Stdout, name, log.LstdFlags)
	l.SetOutput(os.Stdout)

	return &Logger{l, true}
}


