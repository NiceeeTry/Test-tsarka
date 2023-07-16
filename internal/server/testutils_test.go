package server

import (
	"io/ioutil"
	"log"
	"testing"
)

func newTestApplication(t *testing.T) *Application {
	return &Application{
		ErrorLog: log.New(ioutil.Discard, "", 0),
		InfoLog:  log.New(ioutil.Discard, "", 0),
	}
}
