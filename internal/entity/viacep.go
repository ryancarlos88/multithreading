package entity

import "fmt"

type RetornoViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (r *RetornoViaCep) ToString() string {
	if r.Localidade == "" {return "not found"}
	return fmt.Sprintf("Estado: %v, Cidade: %v, Bairro: %v, Rua: %v, CEP: %v", r.Uf, r.Localidade, r.Bairro, r.Logradouro, r.Cep)
}