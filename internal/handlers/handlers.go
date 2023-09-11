package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	"github.com/popooq/rodnoolee_birthday/internal/types"
)

type Handler struct {
	repository storageRepo
}

func New(repository storageRepo) Handler {
	return Handler{
		repository: repository,
	}
}

func (h *Handler) Route() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/insertbday", h.insertBirthday)
	r.Get("/getbday", h.allBirthdays)
	return r
}

func (h *Handler) insertBirthday(w http.ResponseWriter, r *http.Request) {
	var birthday types.Birthday
	err := h.repository.InsertBirthday(birthday)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/plain")

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) allBirthdays(w http.ResponseWriter, r *http.Request) {
	birthdays, err := h.repository.GetAllBirthdays()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("bruh"))
	}
	listOfBirthdays := fmt.Sprintf("%+v", birthdays)

	w.Header().Set("Content-Type", "text/plain")

	w.WriteHeader(http.StatusOK)

	w.Write([]byte(listOfBirthdays))
}
