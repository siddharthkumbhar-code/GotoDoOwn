package main

import (
	"sqlitetutorial.net/go/handlers"
	"net/http"
)

func RegisterRoutes() {

	http.HandleFunc("/addtask", handlers.AddTask)
	http.HandleFunc("/rename", handlers.RenameTask)
	http.HandleFunc("/delete", handlers.DeleteTask)
	http.HandleFunc("/list", handlers.ListAll)
	http.HandleFunc("/status", handlers.ChangeStatus)
	http.HandleFunc("/fetchList", handlers.FetchListByOffsetPagination)
	http.HandleFunc("/fetchData", handlers.FetchListByCursorPagination)
	http.HandleFunc("/addmillionsdata", handlers.AddMillionsData)

}
