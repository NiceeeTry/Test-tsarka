package server

import (
	"net/http"

	"Alikhan.Aitbayev/internal/helpers"
)

func (app *Application) emailHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email string
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	emails := app.emailFinder(input.Email)

	err = app.writeJSON(w, http.StatusOK, helpers.Envelope{"Result": emails}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
