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
		http.Redirect(w, r, "/new", http.StatusMovedPermanently)
		return
	}
}