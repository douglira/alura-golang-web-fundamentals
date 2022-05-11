package main

import (
	"net/http"

	"alura/3-fundamentos-web/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
