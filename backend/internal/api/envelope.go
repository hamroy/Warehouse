// Package api provides the HTTP layer: response envelope, error
// mapping, validation helpers, and handlers.
package api

import (
	"encoding/json"
	"net/http"

	"warehouse-management/backend/internal/item"
)

// Meta carries pagination metadata.
type Meta struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

// Response is the standard success envelope.
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Meta    *Meta  `json:"meta,omitempty"`
}

// ErrorResponse is the standard error envelope.
type ErrorResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Errors  []string `json:"errors,omitempty"`
}

// WriteJSON sends a success envelope.
func WriteJSON(w http.ResponseWriter, status int, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// WritePaginatedJSON sends a success envelope with pagination meta.
func WritePaginatedJSON(w http.ResponseWriter, status int, message string, data any, meta Meta) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    &meta,
	})
}

// WriteError sends an error envelope.
func WriteError(w http.ResponseWriter, status int, message string, errs ...string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(ErrorResponse{
		Success: false,
		Message: message,
		Errors:  errs,
	})
}

// HandleError maps domain errors to the standard error envelope.
// Unknown errors become 500 without leaking internal detail.
func HandleError(w http.ResponseWriter, err error) {
	switch {
	case item.IsValidation(err):
		var vErr *item.ValidationError
		if e, ok := err.(*item.ValidationError); ok {
			vErr = e
		}
		_ = vErr
		WriteError(w, http.StatusBadRequest, "Validation failed", err.Error())
	case item.IsConflict(err):
		WriteError(w, http.StatusConflict, "Conflict", err.Error())
	case item.IsNotFound(err):
		WriteError(w, http.StatusNotFound, "Not found", err.Error())
	default:
		WriteError(w, http.StatusInternalServerError, "Internal server error")
	}
}
