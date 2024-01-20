package employee

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"io"
	"middle-developer-test/internal/app/builder"
	"middle-developer-test/internal/app/model"
	"middle-developer-test/internal/app/options"
	"net/http"
	"strconv"
)

type Handler struct {
	options.HandlerOption
	validator validator.Validate
}

func NewHandler(opt options.HandlerOption, validator validator.Validate) *Handler {
	return &Handler{
		HandlerOption: opt,
		validator:     validator,
	}
}

func (h Handler) GetAll(w http.ResponseWriter, r *http.Request) (data interface{}, err error) {
	ctx := r.Context()
	data, err = h.Services.Employee.GetAll(ctx)
	if err != nil {
		return
	}

	return
}

func (h Handler) GetByID(w http.ResponseWriter, r *http.Request) (data interface{}, err error) {
	ctx := r.Context()
	employeeID, err := strconv.ParseUint(chi.URLParam(r, "employee_id"), 10, 64)
	if err != nil {
		log.Err(err).Send()
		err = builder.WrapError(err, builder.ErrInvalidRequestBody)
		return
	}

	data, err = h.Services.Employee.GetByID(ctx, employeeID)
	if err != nil {
		return
	}

	return
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) (data interface{}, err error) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Err(err).Send()
		err = builder.WrapError(err, builder.ErrInvalidRequestBody)
		return
	}

	var createReq model.EmployeePayload
	err = json.Unmarshal(body, &createReq)
	if err != nil {
		log.Err(err).Send()
		err = builder.WrapError(err, builder.ErrInvalidRequestBody)
		return
	}

	err = h.validator.Struct(createReq)
	if err != nil {
		log.Err(err).Send()
		err = builder.WrapError(err, builder.ErrInvalidRequestBody)
		return
	}

	data, err = h.Services.Employee.Create(ctx, createReq)
	if err != nil {
		return
	}

	return

}

func (h Handler) Update(w http.ResponseWriter, r *http.Request) (data interface{}, err error) {
	ctx := r.Context()
	employeeID, err := strconv.ParseUint(chi.URLParam(r, "employee_id"), 10, 64)
	if err != nil {
		log.Err(err).Send()
		err = builder.WrapError(err, builder.ErrInvalidRequestBody)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Err(err).Send()
		err = builder.WrapError(err, builder.ErrInvalidRequestBody)
		return
	}

	var updateReq model.EmployeePayload
	err = json.Unmarshal(body, &updateReq)
	if err != nil {
		log.Err(err).Send()
		err = builder.WrapError(err, builder.ErrInvalidRequestBody)
		return
	}

	err = h.validator.Struct(updateReq)
	if err != nil {
		log.Err(err).Send()
		err = builder.WrapError(err, builder.ErrInvalidRequestBody)
		return
	}

	updateReq.ID = employeeID
	data, err = h.Services.Employee.Update(ctx, updateReq)
	if err != nil {
		return
	}

	return
}

func (h Handler) Delete(w http.ResponseWriter, r *http.Request) (data interface{}, err error) {
	ctx := r.Context()
	employeeID, err := strconv.ParseUint(chi.URLParam(r, "employee_id"), 10, 64)
	if err != nil {
		log.Err(err).Send()
		err = builder.WrapError(err, builder.ErrInvalidRequestBody)
		return
	}

	err = h.Services.Employee.DeleteByID(ctx, employeeID)
	if err != nil {
		return
	}

	return
}
