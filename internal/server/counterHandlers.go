package server

import (
	"fmt"
	"net/http"

	"Alikhan.Aitbayev/internal/helpers"
)

func (app *Application) addHandler(w http.ResponseWriter, r *http.Request) {
	i, err := app.readParam(r, "i")
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = app.models.Counter.Add(i)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	num, err := app.models.Counter.Get()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, helpers.Envelope{"counter": num}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *Application) subHandler(w http.ResponseWriter, r *http.Request) {
	i, err := app.readParam(r, "i")
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	// ASK IF THE COUNTERS DROPS TO <0
	// should we check if the counter exists?
	err = app.models.Counter.Sub(i)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	num, err := app.models.Counter.Get()
	if err != nil {
		fmt.Println("error")
		return
	}
	err = app.writeJSON(w, http.StatusOK, helpers.Envelope{"counter": num}, nil)
	if err != nil {
		fmt.Println("error")
		return
	}
}

func (app *Application) valHandler(w http.ResponseWriter, r *http.Request) {
	counter, err := app.models.Counter.Get()
	// If the counter is uninitialized, the default value is 0

	// if err != nil {
	// 	fmt.Println("Counter wasnt initialized")
	// 	return
	// }
	err = app.writeJSON(w, http.StatusOK, helpers.Envelope{"counter": counter}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
