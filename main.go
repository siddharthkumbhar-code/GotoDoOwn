package main

import (
	"net/http"

	"sqlitetutorial.net/go/db"
	"sqlitetutorial.net/go/file"
)

func main() {
	file.FileCreation()
	db.DataBaseConnection()
	defer db.DDB.Close()
	RegisterRoutes()
	http.ListenAndServe(":8080", nil)
}
