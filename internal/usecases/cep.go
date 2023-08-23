package usecases

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ryancarlos88/multithreading/config"
	"github.com/ryancarlos88/multithreading/internal/entity"
)

type CepUsecase struct {
	cfg *config.Config
}
func NewCepUsecase(cfg *config.Config) *CepUsecase {
	return &CepUsecase{cfg}
}
func (u *CepUsecase) FetchCepData(cep string) string {
	res := ""
	vcCh := make(chan entity.RetornoViaCep)
	acCh := make(chan entity.RetornoApiCep)

	go u.fetchViaCep(cep, vcCh)
	go u.fetchApiCep(cep, acCh)

	select {
	case r := <- vcCh:
		res = fmt.Sprintf("Retorno da ViaCep: %v", r.ToString())
	case r := <- acCh:
		res = fmt.Sprintf("Retorno da ApiCep: %v", r.ToString())
	case <- time.After(time.Second):
		res = "timeout"
	}

	return res
}

func (u *CepUsecase) fetchViaCep(cep string, ch chan entity.RetornoViaCep) error {
	url := fmt.Sprintf("%v%v%v", u.cfg.ViaCepPrefix, cep, u.cfg.ViaCepSufix)
	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		return  err
	}
	resBytes, _ := io.ReadAll(res.Body)
	var resp entity.RetornoViaCep
	err = json.Unmarshal(resBytes, &resp)
	if err != nil {
		return err
	}

	ch <- resp
	return nil
}
func (u *CepUsecase) fetchApiCep(cep string, ch chan<- entity.RetornoApiCep) error {
	url := fmt.Sprintf("%v%v%v", u.cfg.ApiCepPrefix, cep, u.cfg.ApiCepSuffix)
	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		return  err
	}
	resBytes, _ := io.ReadAll(res.Body)
	var resp entity.RetornoApiCep
	err = json.Unmarshal(resBytes, &resp)
	if err != nil {
		return err
	}

	ch <- resp
	return nil
}
