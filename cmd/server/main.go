package main

import (
	"io"
	"log"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"result": "healthy"}`)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", HealthCheckHandler)
	log.Println("server starting...")
	log.Fatal(http.ListenAndServe(":9595", mux))
}
