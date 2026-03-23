package handlers

import (
	"net/http"

	"sqlitetutorial.net/go/db"
)

func AddMillionsData(w http.ResponseWriter, r *http.Request) {
	query := `INSERT INTO tasks (Name,Status)
			VALUES(?,?)`

	for i := 0; i <= 10000; i++ {
		if i%2 == 0 {
			db.DDB.Exec(query, "learn", true)
		} else {
			db.DDB.Exec(query, "learning", false)
		}
	}
}