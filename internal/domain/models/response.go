package models

type Response struct {
	USDBRL struct {
		Bid  string `json:"bid"`
	} `json:"USDBRL"`
}