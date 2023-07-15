package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"Alikhan.Aitbayev/Data"
	sqlitedb "Alikhan.Aitbayev/Data/sqliteDB"
	"Alikhan.Aitbayev/internal/helpers"
	"github.com/go-redis/redis"
)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	models   Data.Models
}

func NewApplication() (*Application, *sql.DB) {
	InfoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	ErrorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_, err := client.Ping().Result()
	if err != nil {
		ErrorLog.Fatalf("Failed to connect to Redis: %v", err)
	}

	db, err := openDB()
	if err != nil {
		ErrorLog.Fatal(err)
	}

	err = sqlitedb.CreateTables(db)
	if err != nil {
		ErrorLog.Fatal(err)
	}

	return &Application{
		ErrorLog: ErrorLog,
		InfoLog:  InfoLog,
		models:   Data.NewModels(db, client),
	}, db
}

func (app *Application) Run() error {
	srv := &http.Server{
		Addr:     ":8080",
		ErrorLog: app.ErrorLog,
		Handler:  app.routes(),
	}
	app.InfoLog.Println("Starting server on :8080")
	err := srv.ListenAndServe()
	return err
}

func openDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func (app *Application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	return helpers.ReadJSON(w, r, dst)
}

func (app *Application) writeJSON(w http.ResponseWriter, status int, data helpers.Envelope, header http.Header) error {
	return helpers.WriteJSON(w, status, data, header)
}

func (app *Application) longestSubstring(text string) string {
	return helpers.LongestSubstring(text)
}

func (app *Application) emailFinder(emails string) []string {
	return helpers.EmailFinder(emails)
}

func (app *Application) readParam(r *http.Request, paramName string) (int, error) {
	return helpers.ReadParam(r, paramName)
}
