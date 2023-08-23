package entity

import "fmt"

type RetornoApiCep struct {
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

func (r *RetornoApiCep) ToString() string {
	if r.State == "" { return "not found"}
	return fmt.Sprintf("Estado: %v, Cidade: %v, Bairro: %v, Rua: %v, CEP: %v", r.State, r.City, r.District, r.Address, r.Code)
}