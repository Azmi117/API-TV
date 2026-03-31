package apperror

import "net/http"

type Apperror struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Apperror) Error() string {
	return e.Message
}

func NotFound(msg string) error {
	return &Apperror{Code: http.StatusNotFound, Message: msg}
}

func BadRequest(msg string) error {
	return &Apperror{Code: http.StatusBadRequest, Message: msg}
}

func Internal(msg string) error {
	return &Apperror{Code: http.StatusInternalServerError, Message: msg}
}
