package builder

import "net/http"

var (
	ErrInternalServer HttpError = NewHttpError(
		"Ups, something wrong with the server. Please try again later",
		http.StatusInternalServerError,
	)

	ErrInvalidRequestBody HttpError = NewHttpError(
		"Empty or invalid request",
		http.StatusBadRequest,
	)

	ErrEmployeeNotFound HttpError = NewHttpError(
		"Employee data is not found",
		http.StatusNotFound,
	)
)
