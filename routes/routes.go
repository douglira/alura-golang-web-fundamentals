package routes

import (
	"alura/3-fundamentos-web/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.NovoProduto)
	http.HandleFunc("/insert", controllers.Inserir)
	http.HandleFunc("/delete", controllers.Excluir)
	http.HandleFunc("/edit", controllers.Editar)
	http.HandleFunc("/update", controllers.Atualizar)
}
