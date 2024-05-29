package controllers

import (
	"html/template"
	"loja-web-app/db"
	"loja-web-app/produtos"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

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
