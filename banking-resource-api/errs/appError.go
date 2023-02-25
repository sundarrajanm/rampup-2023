package errs

import "net/http"

// swagger:model AppError
type AppError struct {
	// Error code usually maps to the HTTP error code
	// in: int64
	Code int `json:"code"`

	// Message of the error
	// in: string
	Message string `json:"message"`
}

func NewUnexpectedError(msg string) *AppError {
	return &AppError{
		Message: msg,
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(msg string) *AppError {
	return &AppError{
		Message: msg,
		Code:    http.StatusNotFound,
	}
}
