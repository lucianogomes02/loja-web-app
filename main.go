package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"loja-web-app/produtos"
	"net/http"

	_ "github.com/lib/pq"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	// Start the application
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8000", nil)
}

func dbConnection() *sql.DB {
	URI := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"0.0.0.0", 5051, "postgres", "qwerty", "loja")
	db, err := sql.Open("postgres", URI)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	db := dbConnection()

	queryProdutos, err := db.Query("SELECT * FROM produtos")

	if err != nil {
		panic(err.Error())
	}

	produto := produtos.Produto{}
	produtos := []produtos.Produto{}

	for queryProdutos.Next() {
		var id, nome, descricao string
		var preco float64
		var quantidade int

		err := queryProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		produto.ID = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade

		produtos = append(produtos, produto)
	}

	templates.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}
