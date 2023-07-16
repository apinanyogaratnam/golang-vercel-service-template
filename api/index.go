package handler

import (
	"net/http"
	"encoding/json"
)

func Root(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello World")
}

func Handler(w *http.ResponseWriter, r *http.Request) {
	server := http.NewServeMux()

	server.HandleFunc("/", Root)
}
