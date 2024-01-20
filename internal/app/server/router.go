package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"middle-developer-test/internal/app/domain/employee"
	"middle-developer-test/internal/app/options"
)

func Router(opt options.HandlerOption) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Group(func(r chi.Router) {
		employee.Route(r, opt)
	})

	return r
}
