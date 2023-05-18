package utils

import (
	"fmt"
	"io/ioutil"
	"log"

)

func SaveCotacaoToFile(bid string){
	data := []byte(fmt.Sprintf("Dolar: %s", bid))
	if err := ioutil.WriteFile("cotacao.txt", data, 0644); err != nil {
		log.Fatal(err)
	}

}