package builder

import (
	"encoding/json"
	"net/http"
)

type HttpError interface {
	Response() string
	StatusCode() int
}

var _ HttpError = new(httpError)

type httpError struct {
	response   string
	statusCode int
}

func NewHttpError(response string, statusCode int) *httpError {
	return &httpError{
		response:   response,
		statusCode: statusCode,
	}
}

func (h *httpError) Response() (r string) {
	return h.response
}

func (h *httpError) StatusCode() int {
	return h.statusCode
}

type wrappedError struct {
	error
	HttpError
}

func WrapError(err error, httpError HttpError) error {
	return wrappedError{error: err, HttpError: httpError}
}

type HttpHandler struct {
	_ struct{}
	H func(w http.ResponseWriter, r *http.Request) (interface{}, error)
}

func (h HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp, err := h.H(w, r)
	if err != nil {
		errResponse(w, err)
		return
	}

	okResponse(w, resp)
}

func errResponse(w http.ResponseWriter, err error) {
	errResp, ok := err.(HttpError)
	if !ok {
		errResp = ErrInternalServer
	}

	if errResp.Response() == "" && errResp.StatusCode() == 0 {
		errResp = ErrInternalServer
	}

	response := BuildResponse("error", errResp.Response(), errResp.StatusCode())
	writeResponse(w, response, errResp.StatusCode())

}

func okResponse(w http.ResponseWriter, resp interface{}) {
	statusCode := http.StatusOK
	response := BuildResponse("success", resp, statusCode)
	writeResponse(w, response, statusCode)

}

func writeResponse(w http.ResponseWriter, resp interface{}, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	respByte, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ups, something wrong with the server. Please try again later"))
	}

	w.WriteHeader(statusCode)
	w.Write(respByte)
}
