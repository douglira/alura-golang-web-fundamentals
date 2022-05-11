package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaBD() *sql.DB {
	conexao := "user=douglas dbname=alura_loja password=docker sslmode=disable"
	driver := "postgres"
	db, err := sql.Open(driver, conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}
