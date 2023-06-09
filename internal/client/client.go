package client

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/VictorOliveiraPy/internal/domain/models"
	"github.com/VictorOliveiraPy/internal/utils"
)


func GetDollarRate() (*models.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(),   200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET","https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err!= nil {
		log.Println("Erro ao fazer requisição:", err)
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err!= nil {
        log.Println("Erro ao fazer requisição:", err)
        return nil, err
    }
	defer resp.Body.Close()

	var usdbrl  models.Response
	err = json.NewDecoder(resp.Body).Decode(&usdbrl)
	if err != nil {
		log.Println("Erro ao ler json:", err)
		return nil, err
	}

	utils.SaveCotacaoToFile(usdbrl.USDBRL.Bid)

	return &usdbrl, nil

}