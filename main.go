package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

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
	return fmt.Sprintf("Estado: %v, Cidade: %v, Bairro: %v, Rua: %v, CEP: %v", r.State, r.City, r.District, r.Address, r.Code)
}

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
	return fmt.Sprintf("Estado: %v, Cidade: %v, Bairro: %v, Rua: %v, CEP: %v", r.Uf, r.Localidade, r.Bairro, r.Logradouro, r.Cep)
}

func main() {
	apicepCh := make(chan RetornoApiCep)
	viacepCh := make(chan RetornoViaCep)
	cep := "27213-050"
	apicepurl := fmt.Sprintf("https://cdn.apicep.com/file/apicep/%v.json", cep)
	viacepurl := fmt.Sprintf("http://viacep.com.br/ws/%v/json/", cep)

	go func() {
		res, err := http.Get(apicepurl)
		if err != nil {
			panic(err)
		}
		resBytes, _ := io.ReadAll(res.Body)
		var resp RetornoApiCep
		err = json.Unmarshal(resBytes, &resp)
		if err != nil {
			panic(err)
		}

		apicepCh <- resp
	}()

	go func() {
		res, err := http.Get(viacepurl)
		if err != nil {
			panic(err)
		}
		resBytes, _ := io.ReadAll(res.Body)
		var resp RetornoViaCep
		err = json.Unmarshal(resBytes, &resp)
		if err != nil {
			panic(err)
		}

		viacepCh <- resp
	}()

	select {
	case r := <-apicepCh:
		fmt.Printf("Retorno da ApiCep: %v", r.ToString())
	case r := <-viacepCh:
		fmt.Printf("Retorno da ViaCep: %v", r.ToString())
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
}
