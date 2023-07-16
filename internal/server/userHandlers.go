package server

import (
	"database/sql"
	"errors"
	"net/http"

	sqlitedb "Alikhan.Aitbayev/Data/sqliteDB"
	"Alikhan.Aitbayev/internal/helpers"
)

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
