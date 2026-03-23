package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"sqlitetutorial.net/go/db"
	"sqlitetutorial.net/go/file"
	"sqlitetutorial.net/go/models"
)

func FetchListByOffsetPagination(w http.ResponseWriter, r *http.Request) {

	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")
	l, err := strconv.Atoi(limit)
	o, _ := strconv.Atoi(offset)
	if err != nil {
		file.FFile.Write([]byte(err.Error()))
		return
	}

	rows, err := db.DDB.Query(`SELECT Id,Name,Status FROM tasks
			LIMIT ? OFFSET ?`, l, o)
	
	tasks := []models.Task{}
	var task models.Task
	for rows.Next() {
		rows.Scan(&task.Id, &task.Name, &task.Status)
		tasks = append(tasks, task)
	}
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(tasks)
	w.Write(data)
}