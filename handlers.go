package main

import (
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodPost {
	// 	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	// 	return
	// }
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
	// if r.Method != http.MethodPost {
	// 	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	// 	return
	// }
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

func (app *application) addHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodPost {
	// 	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	// 	return
	// }
	// num, err := strconv.Atoi(r.URL.Query().Get("i"))
	// if err != nil {
	// 	http.NotFound(w, r)
	// 	return
	// }
	i, err := app.readIDParam(r)
	if err != nil {
		http.Error(w, "Internal", 500)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", i)
}
