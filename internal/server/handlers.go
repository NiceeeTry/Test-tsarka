package server

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	sqlitedb "Alikhan.Aitbayev/Data/sqliteDB"
	"Alikhan.Aitbayev/internal/helpers"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
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
	//If the counter is uninitialized, the default value is 0

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

func (app *Application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	user := sqlitedb.User{}
	err := app.readJSON(w, r, &user)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	id, err := app.models.Users.Insert(&user)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, helpers.Envelope{"Created user id": id}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *Application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readParam(r, "id")
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	user, err := app.models.Users.Get(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			app.errorResponse(w, r, http.StatusBadRequest, "There is no a user with such id")
		} else {
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	app.writeJSON(w, http.StatusOK, helpers.Envelope{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *Application) putUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readParam(r, "id")
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	user := sqlitedb.User{}
	err = app.readJSON(w, r, &user)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.models.Users.Update(id, &user)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusNoContent, nil, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *Application) deletetUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readParam(r, "id")
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = app.models.Users.Delete(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusNoContent, nil, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
