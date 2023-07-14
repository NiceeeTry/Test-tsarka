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
	i, err := app.readIParam(r)
	if err != nil {
		http.Error(w, "Internal", 500)
		return
	}

	app.dbRedis.IncrBy("counter", i) //check errors
	num, err := app.dbRedis.Get("counter").Int()
	if err != nil {
		fmt.Println("error")
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"counter": num}, nil)
	if err != nil {
		fmt.Println("error")
		return
	}
}

func (app *application) subHandler(w http.ResponseWriter, r *http.Request) {
	i, err := app.readIParam(r)
	if err != nil {
		http.Error(w, "Internal", 500)
		return
	}
	// ASK IF THE COUNTERS DROPS TO <0
	// should we check if the counter exists?

	app.dbRedis.DecrBy("counter", i) //check errors
	num, err := app.dbRedis.Get("counter").Int()
	if err != nil {
		fmt.Println("error")
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"counter": num}, nil)
	if err != nil {
		fmt.Println("error")
		return
	}
}

func (app *application) valHandler(w http.ResponseWriter, r *http.Request) {
	counter, err := app.dbRedis.Get("counter").Int()
	if err != nil {
		fmt.Println("Counter wasnt initialized")
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"counter": counter}, nil)
	if err != nil {
		fmt.Println("error")
		return
	}
}
