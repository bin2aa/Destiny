package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// APIResponse is the standard API response envelope.
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   *APIError   `json:"error,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

// APIError represents an error detail in the response.
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail,omitempty"`
}

// Meta holds pagination metadata.
type Meta struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

// Success sends a 200 OK response with data.
func Success(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Data:    data,
	})
}

// Created sends a 201 Created response.
func Created(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusCreated, APIResponse{
		Success: true,
		Data:    data,
	})
}

// SuccessWithMeta sends a success response with pagination meta.
func SuccessWithMeta(c echo.Context, data interface{}, meta *Meta) error {
	return c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Data:    data,
		Meta:    meta,
	})
}

// Error sends an error response with the given status code and message.
func Error(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, APIResponse{
		Success: false,
		Error: &APIError{
			Code:    statusCode,
			Message: message,
		},
	})
}

// ErrorWithDetail sends an error response with a detail field.
func ErrorWithDetail(c echo.Context, statusCode int, message, detail string) error {
	return c.JSON(statusCode, APIResponse{
		Success: false,
		Error: &APIError{
			Code:    statusCode,
			Message: message,
			Detail:  detail,
		},
	})
}

// Message sends a simple success message response.
func Message(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, APIResponse{
		Success: true,
		Message: message,
	})
}
