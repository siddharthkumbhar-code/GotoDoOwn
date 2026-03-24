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
	Limit, err := strconv.Atoi(limit)
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
	Offset, err:= strconv.Atoi(offset)
	if err!= nil {
		http.Error(w, "Offset must be Valid Integer", 400)
		file.FFile.Write([]byte(err.Error()))
		return
	}
	if Offset<0{
		http.Error(w,"Offset Must Be Greater Than 0",400)
		file.FFile.Write([]byte("Offset Must Be Greater Than 0"))
		return
	}

	rows, err := db.DDB.Query(`SELECT Id,Name,Status FROM tasks
			LIMIT ? OFFSET ?`, Limit, Offset)

	if err!=nil{
		http.Error(w,"Internal Server Error",500)
		file.FFile.Write([]byte(err.Error()))
		return
	}
	tasks := []models.Task{}
	for rows.Next() {
		var task models.Task
		err:=rows.Scan(&task.Id, &task.Name, &task.Status)
		if err != nil {
   		 	http.Error(w,"Internal Server Error",500)
    	 	return
	    }
		tasks = append(tasks, task)
	}
	defer rows.Close()
	nextOffset := Offset + Limit

	response := map[string]interface{}{
    	"data": tasks,
    	"next_offset": nextOffset,
    	"has_more": len(tasks) == Limit,
	}
	w.Header().Set("Content-Type", "application/json")
	responseData, err := json.Marshal(response)
	if err!=nil{
		http.Error(w,"Internal Server Error",500)
		file.FFile.Write([]byte("Internal Server Error"))
		return
	}
	w.Write(responseData)
}