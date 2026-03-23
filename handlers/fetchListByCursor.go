package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"sqlitetutorial.net/go/db"
	"sqlitetutorial.net/go/models"
)

func FetchListByCursorPagination(w http.ResponseWriter, r *http.Request) {

	cid := r.URL.Query().Get("cursor")
	limit := r.URL.Query().Get("limit")
	ciid, _ := strconv.Atoi(cid)
	l, _ := strconv.Atoi(limit)

	query := `SELECT * FROM tasks
			WHERE Id>?
			ORDER BY Id
			LIMIT ? `
	rows, err := db.DDB.Query(query, ciid, l)
	if err != nil {
		http.Error(w, "Database Error", 500)
		return
	}

	var task models.Task
	tasks := []models.Task{}

	for rows.Next() {
		rows.Scan(&task.Id, &task.Name, &task.Status)
		tasks = append(tasks, task)
	}

	w.Header().Set("Content-type", "application/json")
	data, err := json.Marshal(tasks)
	w.Write(data)
}