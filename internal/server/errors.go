package server

import (
	"fmt"
	"net/http"

	"Alikhan.Aitbayev/internal/helpers"
)

func (app *Application) logError(r *http.Request, err error) {
	app.ErrorLog.Println(err, map[string]string{
		"request_method": r.Method,
		"request_url":    r.URL.String(),
	})
}

func (app *Application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := helpers.Envelope{
		"error": message,
	}
	err := app.writeJSON(w, status, env, nil)
	app.logError(r, err)
	if err != nil {
		w.WriteHeader(500)
	}
}

func (app *Application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *Application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

func (app *Application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func (app *Application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}
