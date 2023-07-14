package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	sqlitedb "Alikhan.Aitbayev/internal/sqliteDB"
	"github.com/go-redis/redis"
	_ "github.com/mattn/go-sqlite3"
)

// var client *redis.Client //Check on correctnes
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	dbRedis  *redis.Client
	models   sqlitedb.Models
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	db, err := openDB()
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	err = createTables(db)
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		dbRedis:  client,
		models:   sqlitedb.NewModels(db),
	}

	srv := &http.Server{
		Addr:     ":8080",
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	infoLog.Println("Starting server on :8080")
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

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
