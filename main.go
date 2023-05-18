package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" //driver
	"github.com/VictorOliveiraPy/internal/repository"


)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal("Erro ao abrir o banco de dados:", err)
	}
	defer db.Close()

	err = repository.CreateTable(db)
	if err != nil {
		log.Fatal("Erro ao criar a tabela:", err)
	}
}
