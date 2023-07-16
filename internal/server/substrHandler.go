package server

import (
	"net/http"

	"Alikhan.Aitbayev/internal/helpers"
)

func (app *Application) substrHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Text string `json:"text"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	longestSubstring := app.longestSubstring(input.Text)
	err = app.writeJSON(w, http.StatusOK, helpers.Envelope{"Result": longestSubstring}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
