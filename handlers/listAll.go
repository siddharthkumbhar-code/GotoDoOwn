package handlers
import (
	"encoding/json"
	"net/http"

	"sqlitetutorial.net/go/db"
	"sqlitetutorial.net/go/file"
	"sqlitetutorial.net/go/models"
)


func ListAll(w http.ResponseWriter, r *http.Request) {
	query := `SELECT Id,Name,Status FROM tasks`
	rows, err := db.DDB.Query(query)
	if err != nil {
		file.FFile.Write([]byte(err.Error()))
		http.Error(w, err.Error(), 400)
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