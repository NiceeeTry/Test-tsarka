package main

import (
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var input struct {
		Text string `json:"text"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// delete after
	comments, err := app.dbRedis.LRange("comments", 0, 10).Result()
	fmt.Println(comments)

	longestSubstring := app.LongestSubstring(input.Text)
	err = app.writeJSON(w, http.StatusOK, envelope{"Result": longestSubstring}, nil)
}

func (app *application) emailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var input struct {
		Email string
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	emails := app.emailFinder(input.Email)

	err = app.writeJSON(w, http.StatusOK, envelope{"Result": emails}, nil)
}
