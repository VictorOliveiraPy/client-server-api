package repository

import (
	"database/sql"
	"github.com/VictorOliveiraPy/internal/domain/models"
	_ "github.com/mattn/go-sqlite3" // Importar o driver SQLite3

)


func CreateTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS cotacoes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		bid TEXT
	)`)
	return err
}

func NewDataRepository(db *sql.DB, response *models.Response) error{
	query := "INSERT INTO cotacoes (bid) VALUES (?)"
	stmt, err := db.Prepare(query)
	if err!= nil {
        return err
    }
	
	defer stmt.Close()

	_, err = stmt.Exec(response.USDBRL.Bid)
	if err != nil {
		return err
	}

	return nil
	
}