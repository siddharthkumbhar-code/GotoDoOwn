package handlers

import (
	//"GoToDo/models"
	"encoding/json"
	"net/http"
	"strconv"

	"sqlitetutorial.net/go/db"
	"sqlitetutorial.net/go/models"
	// "GoToDo/db"
	//"go.uber.org/mock/mockgen/model"
)


func RenameTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	iid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "error", 400)
		return
	}
	var task models.Task
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "error", 400)
		return
	}
	query := `UPDATE tasks 
			SET Name=?
			WHERE id=?`
	_,err=db.DDB.Exec(query, task.Name, iid)
	if err!=nil{
		http.Error(w,"INternal Server Error",500)
		return
	}
	w.Write([]byte("Rename Successfully Done"))
}