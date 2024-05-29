package main

import (
	"html/template"
	"loja-web-app/produtos"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	// Start the application
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8000", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	produtos := []produtos.Produto{
		{ID: 1, Nome: "Camiseta", Descricao: "Camiseta preta", Preco: 39.90, Quantidade: 300},
		{ID: 2, Nome: "Calça", Descricao: "Calça jeans", Preco: 89.90, Quantidade: 100},
		{ID: 3, Nome: "Tênis", Descricao: "Tênis esportivo", Preco: 159.90, Quantidade: 50},
	}
	templates.ExecuteTemplate(w, "Index", produtos)
}
