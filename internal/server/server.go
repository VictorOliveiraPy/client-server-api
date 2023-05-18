package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/VictorOliveiraPy/internal/client"
	"github.com/VictorOliveiraPy/internal/repository"

	_ "github.com/mattn/go-sqlite3" //driver
)

func main(){
	http.HandleFunc("/cotacao", handler)
	http.ListenAndServe(":8080", nil)

}


func handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	resp, err := client.GetDollarRate()
	if err != nil {
        log.Println("Erro ao obter a cotação do dólar:", err)
		return
	}

	select {
		case <- time.After(10 * time.Second):
			log.Println("Tempo esgostado para salvar no banco...")
			http.Error(w, "Tempo esgostado para salvar no banco...", http.StatusRequestTimeout)

		case <-ctx.Done():
			log.Println("Request canceled for the client")
			http.Error(w, "Request canceled for the client", http.StatusRequestTimeout)

		default:
			db, err := sql.Open("sqlite3", "test.db")
			if err != nil {
				log.Fatal("Erro ao abrir o banco de dados:", err)
			}
			defer db.Close()
			log.Println(resp)
			err = repository.NewDataRepository(db, resp)
			if err!= nil {
                log.Fatal("Erro ao salvar no banco de dados:", err)
            }
			w.WriteHeader(http.StatusCreated)
			fmt.Fprint(w, "Cotação salva no banco de dados")
			log.Println("Cotacao Salvar com sucesso!")

	}
}