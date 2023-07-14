package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/rest/substr/find", app.home)
	router.HandlerFunc(http.MethodPost, "/rest/email/check", app.emailHandler)
	router.HandlerFunc(http.MethodPost, "/rest/counter/add/:i", app.addHandler)
	// mux := http.NewServeMux()
	// mux.HandleFunc("/rest/substr/find", app.home)
	// mux.HandleFunc("/rest/email/check", app.emailHandler)
	// mux.HandleFunc("/rest/counter/add/", app.addHandler)
	// mux.HandleFunc("/rest/counter/sub/", app.subHandler)
	// mux.HandleFunc("/rest/counter/val", app.valHandler)
	return router
}
