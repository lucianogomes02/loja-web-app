package controllers

import (
	"html/template"
	"loja-web-app/db"
	"loja-web-app/produtos"
	"net/http"
	"strconv"
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

func NewProductHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, erro := strconv.ParseFloat(preco, 64)

		if erro != nil {
			panic(erro.Error())
		}

		quantidadeConvertida, erro := strconv.Atoi(quantidade)

		if erro != nil {
			panic(erro.Error())
		}

		repository := produtos.NewProdutoRepository(
			db.DbConnection(),
		)
		repository.CriaProduto(nome, descricao, precoConvertido, quantidadeConvertida)
	}

	http.Redirect(w, r, "/", 301)
}
