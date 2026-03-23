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
	if limit==""{
		http.Error(w,"Limit Is Required",400)
		file.FFile.Write([]byte("Limit is Required"))
		return
	}
	offset := r.URL.Query().Get("offset")
	if offset==""{
		http.Error(w,"Offset is required",400)
		file.FFile.Write([]byte("Offset is Required"))
		return
	}
	l, err := strconv.Atoi(limit)
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
	o, _ := strconv.Atoi(offset)
	if err != nil {
		http.Error(w, "Offset must be Valid Integer", 400)
		file.FFile.Write([]byte(err.Error()))
		return
	}
	if o<=0{
		http.Error(w,"Offset Must Be Greater Than 0",400)
		file.FFile.Write([]byte("Offset Must Be Greater Than 0"))
		return
	}

	rows, err := db.DDB.Query(`SELECT Id,Name,Status FROM tasks
			LIMIT ? OFFSET ?`, l, o)

	if err!=nil{
		http.Error(w,"Internal Server Error",500)
		file.FFile.Write([]byte(err.Error()))
		return
	}
	
	tasks := []models.Task{}
	var task models.Task
	for rows.Next() {
		rows.Scan(&task.Id, &task.Name, &task.Status)
		tasks = append(tasks, task)
	}
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(tasks)
	if err!=nil{
		http.Error(w,"Internal Server Error",500)
		file.FFile.Write([]byte("Internal Server Error"))
		return
	}
	w.Write(data)
}