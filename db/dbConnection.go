package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"sqlitetutorial.net/go/file"
)

var DDB *sql.DB
func DataBaseConnection() {
	var err error
	DDB, err = sql.Open("sqlite3", "database.db")
	if err != nil {
		//fmt.Println(err)
		file.FFile.Write([]byte(err.Error()))
		return
	}
	//defer db.Close()
	createTable(DDB)
}
func createTable(DDB *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS tasks(
			Id INTEGER PRIMARY KEY AUTOINCREMENT ,
			Name TEXT NOT NULL,
			Status bool NOT NULL);`
	_, err := DDB.Exec(query)
	if err != nil {
		file.FFile.Write([]byte(err.Error()))
		return
	}
}