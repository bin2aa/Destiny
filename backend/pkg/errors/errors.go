package errors

import (
	"net/http"
)

// AppError represents a structured API error.
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail,omitempty"`
}

func (e *AppError) Error() string {
	return e.Message
}

// Common error constructors
func NewNotFoundError(resource string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: resource + " not found",
	}
}

func NewBadRequestError(message string) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewUnauthorizedError(message string) *AppError {
	if message == "" {
		message = "unauthorized"
	}
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}

func NewForbiddenError(message string) *AppError {
	if message == "" {
		message = "forbidden"
	}
	return &AppError{
		Code:    http.StatusForbidden,
		Message: message,
	}
}

func NewInternalError(message string) *AppError {
	if message == "" {
		message = "internal server error"
	}
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func NewConflictError(message string) *AppError {
	return &AppError{
		Code:    http.StatusConflict,
		Message: message,
	}
}

func NewValidationError(detail string) *AppError {
	return &AppError{
		Code:    http.StatusUnprocessableEntity,
		Message: "validation failed",
		Detail:  detail,
	}
}
