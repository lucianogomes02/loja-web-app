package routes

import (
	"loja-web-app/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.IndexHandler)
	http.HandleFunc("/new", controllers.NewProductHandler)
	http.HandleFunc("/insert", controllers.CreateProductHandler)
}
