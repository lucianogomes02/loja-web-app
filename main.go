package main

import (
	"loja-web-app/routes"
	"net/http"
)

func main() {
	// Start the application
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
