package main

import (
	"backend-worktask/tasks"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /task/add", tasks.PostAddHandler)
	mux.HandleFunc("GET /task/drive/", tasks.GetDriveHandler)
}
