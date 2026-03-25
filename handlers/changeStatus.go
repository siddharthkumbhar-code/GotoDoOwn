package handlers

import (
	"net/http"
	"strconv"

	"sqlitetutorial.net/go/db"
	"sqlitetutorial.net/go/file"
)

func ChangeStatus(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		http.Error(w, "Invalid Method", 405)
		file.FFile.Write([]byte("Invalid Method"))
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Id is Required", 400)
		file.FFile.Write([]byte("Id is Required"))
		return
	}
	iid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Id must be Valid Integer", 400)
		file.FFile.Write([]byte(err.Error()))
		return
	}
	if iid <= 0 {
		http.Error(w, "Id Must Be Greater Than 0", 400)
		file.FFile.Write([]byte("Id Must Be Greater Than 0"))
		return
	}

	query := `UPDATE tasks 
			SET Status=1-Status
			WHERE Id=?`

	_, err = db.DDB.Exec(query, iid)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		file.FFile.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Status Change Successfully"))
}
