package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"Alikhan.Aitbayev/Data"
	"github.com/go-redis/redis"
)

type application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	models   Data.Models
}

func NewApplication() (*application, *sql.DB) {
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

	err = createTables(db)
	if err != nil {
		ErrorLog.Fatal(err)
	}

	return &application{
		ErrorLog: ErrorLog,
		InfoLog:  InfoLog,
		models:   Data.NewModels(db, client),
	}, db
}

func (app *application) Run() error {
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

func createTables(db *sql.DB) error {
	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS users 
	(id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		surname TEXT NOT NULL);`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
