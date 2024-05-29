package routes

import (
	"loja-web-app/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.IndexHandler)
}
