package views

import (
	"net/http"

	"github.com/apinanyogaratnam/golang-vercel-service-template/api/controllers"
)

func View(w http.ResponseWriter, r *http.Request) {
	server := http.NewServeMux()

	server.HandleFunc("/", controllers.Root)
	server.HandleFunc("/health", controllers.Health)

	server.ServeHTTP(w, r)
}
