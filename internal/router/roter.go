package router

import (
	"Di-Test/internal/handler"

	"github.com/go-chi/chi/v5"
)

func Router(h *handler.UserHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", h.CreateUser)
	return r
}
