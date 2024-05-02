package response_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brxyxn/go-phonebook-api/internal/response"
	"github.com/stretchr/testify/assert"
)

func TestSuccessResponseReturnsExpectedOutput(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	data := "test data"
	message := "test message"

	response.Success(w, r, data, message)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), data)
	assert.Contains(t, w.Body.String(), message)
}

func TestCreatedResponseReturnsExpectedOutput(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	data := "test data"
	message := "test message"

	response.Created(w, r, data, message)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), data)
	assert.Contains(t, w.Body.String(), message)
}

func TestDeletedResponseReturnsExpectedOutput(t *testing.T) {
	w := httptest.NewRecorder()

	response.Deleted(w)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestErrorResponseReturnsExpectedOutput(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	errorMessage := "test error"

	response.Error(w, r, errorMessage)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), errorMessage)
}

func TestNotFoundResponseReturnsExpectedOutput(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	response.NotFound(w, r)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), response.ErrorNotFound)
}

func TestInternalErrorResponseReturnsExpectedOutput(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	data := "test data"

	response.InternalError(w, r, data)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), data)
}
