package handler

import (
	"net/http"

	"github.com/apinanyogaratnam/golang-vercel-service-template/api/views"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	views.View(w, r)
}
