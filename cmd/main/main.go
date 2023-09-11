package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/popooq/rodnoolee_birthday/internal/config"
	"github.com/popooq/rodnoolee_birthday/internal/handlers"
	"github.com/popooq/rodnoolee_birthday/internal/repository/mongorepo"
)

func main() {
	ctx := context.Background()
	cfg := config.New()
	repo, err := mongorepo.New(ctx, cfg.DBAddress)
	if err != nil {
		log.Fatal(err)
	}

	hndlr := handlers.New(repo)

	rt := chi.NewRouter()
	rt.Mount("/", hndlr.Route())

	log.Fatal(http.ListenAndServe(cfg.Address, rt))
}
