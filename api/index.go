package handler

import (
	"net/http"

	"github.com/apinanyogaratnam/golang-vercel-service-template/views"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	views.Handler(w, r)
}
