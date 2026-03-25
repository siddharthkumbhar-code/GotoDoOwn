package handlers

import (
	"encoding/json"
	"net/http"
	"sqlitetutorial.net/go/db"
	"sqlitetutorial.net/go/file"
	"sqlitetutorial.net/go/models"
)


func AddTask(w http.ResponseWriter, r *http.Request) {

	if r.Method!=http.MethodPost{
		http.Error(w,"Method is Invalid",405)
		http.Error(w,"Method is Invalid",405)
		file.FFile.Write([]byte("Invalid Method Type"))
		return
	}

	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "error", 400)
		file.FFile.Write([]byte(err.Error()))
		return
	}
	query := `INSERT INTO tasks(Name,Status)
	        VALUES(?,?)`
	_, err =db.DDB.Exec(query, task.Name, task.Status)
	if err != nil {
		http.Error(w, err.Error(), 400)
		file.FFile.Write([]byte(err.Error()))
		return
	}
}