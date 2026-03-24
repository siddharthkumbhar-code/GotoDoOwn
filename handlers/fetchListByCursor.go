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
	cursor, err:= strconv.Atoi(cid)
	if err != nil {
		http.Error(w, "Cursor must be Valid Integer", 400)
		file.FFile.Write([]byte(err.Error()))
		return
	}
	if cursor<=0{
		http.Error(w,"Cursor Must Be Greater Than 0",400)
		file.FFile.Write([]byte("Cursor Must Be Greater Than 0"))
		return
	}
	Limit, err:= strconv.Atoi(limit)
	if err != nil {
		http.Error(w, "Limit must be Valid Integer", 400)
		file.FFile.Write([]byte(err.Error()))
		return
	}
	if Limit<=0{
		http.Error(w,"Limit Must Be Greater Than 0",400)
		file.FFile.Write([]byte("Limit Must Be Greater Than 0"))
		return
	}

	query := `SELECT * FROM tasks
			WHERE Id>?
			ORDER BY Id
			LIMIT ? `

	rows, err := db.DDB.Query(query, cursor, Limit)
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

	nextCursor:=cursor+Limit

	response:=map[string]interface{}{
		"data":tasks,
		"next_Cursor":nextCursor,
		"has_more":len(tasks)==Limit,
	}

	w.Header().Set("Content-type", "application/json")
	responsedata, err := json.Marshal(response)
	if err!=nil{
		http.Error(w,"Internal Server Error",500)
		file.FFile.Write([]byte("Internal Server Error"))
		return
	}
	w.Write(responsedata)
}