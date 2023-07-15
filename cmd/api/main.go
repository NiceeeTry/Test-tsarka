package main

import (
	"Alikhan.Aitbayev/internal/server"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app, db := server.NewApplication()
	defer db.Close()

	err := app.Run()
	app.ErrorLog.Fatal(err)

}
