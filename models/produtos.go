package models

import (
	db "alura/3-fundamentos-web/database"
)

type Produto struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

func ListaTodosProdutos() []Produto {
	db := db.ConectaBD()
	query, err := db.Query("SELECT * FROM produtos")

	if err != nil {
		panic(err.Error())
	}

	produtos := []Produto{}

	for query.Next() {
		p := Produto{}

		err := query.Scan(&p.Id, &p.Nome, &p.Preco, &p.Quantidade, &p.Descricao)
		if err != nil {
			panic(err.Error())
		}

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func InserirNovoProduto(nome string, preco float64, descricao string, quantidade int) {
	db := db.ConectaBD()
	query, err := db.Prepare("INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	_, err = query.Exec(nome, descricao, preco, quantidade)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
