package views

import (
	"net/http"

	"github.com/apinanyogaratnam/golang-vercel-service-template/api/controllers"
)

func View() {
	server := http.NewServeMux()

	server.HandleFunc("/", controllers.Root)
}
