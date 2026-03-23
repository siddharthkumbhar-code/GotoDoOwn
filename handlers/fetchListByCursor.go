package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"sqlitetutorial.net/go/db"
	"sqlitetutorial.net/go/file"
	"sqlitetutorial.net/go/models"
)

func FetchListByCursorPagination(w http.ResponseWriter, r *http.Request) {

	cid := r.URL.Query().Get("cursor")
	if cid==""{
		http.Error(w,"cursor Is Required",400)
		file.FFile.Write([]byte("Limit is Required"))
		return
	}
	limit := r.URL.Query().Get("limit")
	if limit==""{
		http.Error(w,"Limit Is Required",400)
		file.FFile.Write([]byte("Limit is Required"))
		return
	}
	ciid, err:= strconv.Atoi(cid)
	if err != nil {
		http.Error(w, "Cursor must be Valid Integer", 400)
		file.FFile.Write([]byte(err.Error()))
		return
	}
	if ciid<=0{
		http.Error(w,"Cursor Must Be Greater Than 0",400)
		file.FFile.Write([]byte("Cursor Must Be Greater Than 0"))
		return
	}
	l, err:= strconv.Atoi(limit)
	if err != nil {
		http.Error(w, "Limit must be Valid Integer", 400)
		file.FFile.Write([]byte(err.Error()))
		return
	}
	if l<=0{
		http.Error(w,"Limit Must Be Greater Than 0",400)
		file.FFile.Write([]byte("Limit Must Be Greater Than 0"))
		return
	}

	query := `SELECT * FROM tasks
			WHERE Id>?
			ORDER BY Id
			LIMIT ? `
	rows, err := db.DDB.Query(query, ciid, l)
	if err != nil {
		http.Error(w, "Database Error", 500)
		file.FFile.Write([]byte(err.Error()))
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
	if err!=nil{
		http.Error(w,"Internal Server Error",500)
		file.FFile.Write([]byte("Internal Server Error"))
		return
	}
	w.Write(data)
}