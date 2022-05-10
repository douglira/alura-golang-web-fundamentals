package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

func conectaBD() *sql.DB {
	conexao := "user=douglas dbname=alura_loja password=docker sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}

type Produto struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectaBD()
	queryProdutos, err := db.Query("SELECT * FROM produtos")

	if err != nil {
		panic(err.Error())
	}

	produtos := []Produto{}

	for queryProdutos.Next() {
		p := Produto{}

		err := queryProdutos.Scan(&p.Id, &p.Nome, &p.Preco, &p.Quantidade, &p.Descricao)
		if err != nil {
			panic(err.Error())
		}

		produtos = append(produtos, p)
	}

	templates.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}
