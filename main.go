package main

import (
	"html/template"
	"loja-web-app/db"
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
	repository := produtos.NewProdutoRepository(
		db.DbConnection(),
	)
	produtos, erro := repository.BuscaTodosProdutos()

	if erro != nil {
		panic(erro.Error())
	}

	templates.ExecuteTemplate(w, "Index", produtos)
}
