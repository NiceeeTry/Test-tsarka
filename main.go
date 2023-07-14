package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
)

// var client *redis.Client //Check on correctnes
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	dbRedis  *redis.Client
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		dbRedis:  client,
	}

	srv := &http.Server{
		Addr:     ":8080",
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	infoLog.Println("Starting server on :8080")
	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}
