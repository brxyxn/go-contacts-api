// Package contacts provides the handlers for the contacts API.
package contacts

import "github.com/brxyxn/go-phonebook-api/pkg/logger"

// Handler is the struct that holds the logger used by the contacts handlers.
type Handler struct {
	// logger is the logger used by the contacts handlers.
	logger logger.SimpleLogger
}

// NewHandler creates a new Handler with the provided logger.
// It returns a pointer to the created Handler.
func NewHandler(logger *logger.SimpleLogger) *Handler {
	return &Handler{*logger}
}
