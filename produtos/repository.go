package produtos

import (
	"database/sql"
)

type ProdutoRepository struct {
	db *sql.DB
}

func NewProdutoRepository(db *sql.DB) *ProdutoRepository {
	return &ProdutoRepository{db: db}
}

func (produtoRepository *ProdutoRepository) BuscaTodosProdutos() ([]Produto, error) {
	queryProdutos, err := produtoRepository.db.Query("SELECT * FROM produtos")

	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}
	produtos := []Produto{}

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
	defer produtoRepository.db.Close()

	return produtos, err
}

func (produtoRepository *ProdutoRepository) CriaProduto(nome, descricao string, preco float64, quantidade int) {
	insertProduto, err := produtoRepository.db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertProduto.Exec(nome, descricao, preco, quantidade)
	defer produtoRepository.db.Close()
}
