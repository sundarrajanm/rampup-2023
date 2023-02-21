package errs

import "net/http"

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewUnexpectedError(msg string) *AppError {
	return &AppError{
		Message: msg,
		Code:    http.StatusInternalServerError,
	}
}
