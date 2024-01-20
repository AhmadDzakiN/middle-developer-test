package employee

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"middle-developer-test/internal/app/builder"
	"middle-developer-test/internal/app/options"
	"net/http"
)

func Route(r chi.Router, opt options.HandlerOption) {
	validator := validator.New()
	employeeHandler := NewHandler(opt, *validator)
	r.Route("/employees", func(r chi.Router) {
		r.Method(http.MethodGet, "/", builder.HttpHandler{H: employeeHandler.GetAll})
		r.Method(http.MethodGet, "/{employee_id}", builder.HttpHandler{H: employeeHandler.GetByID})
		r.Method(http.MethodPost, "/", builder.HttpHandler{H: employeeHandler.Create})
		r.Method(http.MethodPut, "/{employee_id}", builder.HttpHandler{H: employeeHandler.Update})
		r.Method(http.MethodDelete, "/{employee_id}", builder.HttpHandler{H: employeeHandler.Delete})
	})
}
