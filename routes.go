package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/rest/substr/find", app.home)
	mux.HandleFunc("/rest/email/check", app.emailHandler)
	return mux
}
