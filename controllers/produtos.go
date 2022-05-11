package controllers

import (
	"alura/3-fundamentos-web/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.ListaTodosProdutos()
	templates.ExecuteTemplate(w, "Index", produtos)
}

func NovoProduto(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Inserir(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertidoFloat, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantidadeConvertidaInt, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.InserirNovoProduto(nome, precoConvertidoFloat, descricao, quantidadeConvertidaInt)
	}
	http.Redirect(w, r, "/new", http.StatusMovedPermanently)
}

func Excluir(w http.ResponseWriter, r *http.Request) {
	idProduto, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		log.Println("Erro ao obter o ID do produto:", err)
	}

	models.Deletar(idProduto)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Editar(w http.ResponseWriter, r *http.Request) {
	idProduto, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		log.Println("Erro ao obter o ID do produto:", err)
	}

	produto := models.BuscarPorId(idProduto)
	templates.ExecuteTemplate(w, "Edit", produto)
}

func Atualizar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertidoFloat, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		idConvertidoInt, err := strconv.Atoi(id)

		if err != nil {
			log.Println("Erro na conversão do ID:", err)
		}

		quantidadeConvertidaInt, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.AtualizarProduto(idConvertidoInt, nome, descricao, precoConvertidoFloat, quantidadeConvertidaInt)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
