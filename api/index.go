package handler

import (
	"encoding/json"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
}

func Handler(w *http.ResponseWriter, r *http.Request) {
	server := http.NewServeMux()

	server.HandleFunc("/", Root)
}
