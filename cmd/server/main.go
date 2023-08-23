package main

import (
	"net/http"

	"github.com/ryancarlos88/multithreading/config"
	"github.com/ryancarlos88/multithreading/internal/infra/webserver/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main(){
	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	h := handlers.NewHandler(cfg)

	r.Get("/cep", h.CepHandler)
	http.ListenAndServe(":8080", r)
}