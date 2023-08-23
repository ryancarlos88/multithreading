package handlers

import (
	"net/http"

	"github.com/ryancarlos88/multithreading/config"
	"github.com/ryancarlos88/multithreading/internal/usecases"
)

type Handler struct {
	cfg *config.Config
}
func NewHandler(cfg *config.Config) *Handler{
	return &Handler{cfg}
}

func (h *Handler) CepHandler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	uc := usecases.NewCepUsecase(h.cfg)

	res := uc.FetchCepData(cep)

	w.Write([]byte(res))
}
