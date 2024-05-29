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
