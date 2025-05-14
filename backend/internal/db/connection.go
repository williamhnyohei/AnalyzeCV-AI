package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	connStr := "postgres://admin:admin@localhost:5432/analyzecvai?sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar no PostgreSQL:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("PostgreSQL indisponível:", err)
	}

	fmt.Println("Conexão com PostgreSQL estabelecida com sucesso")
}
