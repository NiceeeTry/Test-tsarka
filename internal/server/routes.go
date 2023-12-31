package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/rest/substr/find", app.substrHandler)

	router.HandlerFunc(http.MethodPost, "/rest/email/check", app.emailHandler)

	router.HandlerFunc(http.MethodPost, "/rest/counter/add/:i", app.addHandler)
	router.HandlerFunc(http.MethodPost, "/rest/counter/sub/:i", app.subHandler)
	router.HandlerFunc(http.MethodGet, "/rest/counter/val", app.valHandler)

	router.HandlerFunc(http.MethodPost, "/rest/user", app.registerUserHandler)
	router.HandlerFunc(http.MethodGet, "/rest/user/:id", app.getUserHandler)
	router.HandlerFunc(http.MethodPut, "/rest/user/:id", app.putUserHandler)
	router.HandlerFunc(http.MethodDelete, "/rest/user/:id", app.deletetUserHandler)

	return router
}
