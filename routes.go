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
	router.HandlerFunc(http.MethodPost, "/rest/counter/sub/:i", app.subHandler)
	router.HandlerFunc(http.MethodPost, "/rest/counter/val", app.valHandler)
	return router
}
