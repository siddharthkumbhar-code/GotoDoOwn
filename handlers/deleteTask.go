package handlers

import (
	"net/http"
	"strconv"

	"sqlitetutorial.net/go/db"
	"sqlitetutorial.net/go/file"
)

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id==""{
		http.Error(w,"Id is Required",400)
		file.FFile.Write([]byte("Id is Required"))
		return
	}
	iid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Id must be Valid Integer", 400)
		file.FFile.Write([]byte(err.Error()))
		return
	}
	if iid<=0{
		http.Error(w,"Id Must Be Greater Than 0",400)
		file.FFile.Write([]byte("Id Must Be Greater Than 0"))
		return
	}
	query := `DELETE FROM tasks
			WHERE Id=?`
	_,err=db.DDB.Exec(query, iid)
	if err!=nil{
		http.Error(w,"Internal Server Error",400)
		file.FFile.Write([]byte("Internal Server Error"))
		return
	}
}