package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/apinanyogaratnam/golang-vercel-service-template/models"
)

func Root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.RootResponse{Status: "OK"})
}
