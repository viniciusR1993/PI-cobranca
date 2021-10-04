package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

//Função para Conectar ao Banco de Dados Postgres
func ConectBD() *sql.DB {
	connectino := "user=postgres dbname=fatec password=44432751 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connectino)
	if err != nil {
		panic(err.Error())
	}
	return db
}
