package bootstrap

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/popooq/rodnoolee_birthday/internal/config"
	"github.com/popooq/rodnoolee_birthday/internal/handler"
	"github.com/popooq/rodnoolee_birthday/internal/repository/mongorepo"
	"github.com/popooq/rodnoolee_birthday/internal/usecase/addbirthday"
	"github.com/popooq/rodnoolee_birthday/internal/usecase/allbirthdays"
	"github.com/popooq/rodnoolee_birthday/internal/usecase/updatebirthday"
)

func Bootstrap() error {
	ctx := context.Background()
	cfg := config.New()
	repo, err := mongorepo.New(ctx, cfg.DBAddress)
	if err != nil {
		return err
	}

	hndlr := handler.New(repo)

	rt := chi.NewRouter()
	rt.Mount("/", Route())

	return http.ListenAndServe(cfg.Address, rt)
}

func Route() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/insertbday", addbirthday.InsertBirthday)
	r.Get("/getbday", allbirthdays.AllBirthdays)
	r.Post("/updatebday", updatebirthday.UpdateBirthday)
	return r
}
