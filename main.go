package main

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	// Start the application
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8000", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}
