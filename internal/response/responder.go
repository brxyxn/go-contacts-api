package response

import (
	"encoding/json"
	"github.com/brxyxn/go-phonebook-api/pkg/logger"
	"net/http"
)

const (
	ErrorNotFound       = "data not found"
	ErrorInternalServer = "internal server error"
)

const (
	ActionContinue = "CONTINUE"
	ActionStop     = "STOP"
	ActionError    = "ERROR"
)

type Response[T any] struct {
	Data    T      `json:"data"`
	Message string `json:"message,omitempty"`
	Action  string `json:"action,omitempty"`
}

func Success(w http.ResponseWriter, r *http.Request, data interface{}, message string) {
	response := Response[any]{
		Data:    data,
		Message: message,
		Action:  ActionContinue,
	}

	res, err := json.Marshal(response)
	if err != nil {
		InternalError(w, r, ErrorInternalServer)
	}

	w = SetContentType(w, r, http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		InternalError(w, r, ErrorInternalServer)
	}
}

func Created(w http.ResponseWriter, r *http.Request, data interface{}, message string) {
	response := Response[any]{
		Data:    data,
		Message: message,
		Action:  ActionContinue,
	}

	res, err := json.Marshal(response)
	if err != nil {
		InternalError(w, r, ErrorInternalServer)
	}

	w = SetContentType(w, r, http.StatusCreated)
	_, err = w.Write(res)
	if err != nil {
		InternalError(w, r, ErrorInternalServer)
	}
}

func Deleted(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func Error(w http.ResponseWriter, r *http.Request, error string) {
	response := Response[any]{
		Data:    nil,
		Message: error,
		Action:  ActionError,
	}
	res, err := json.Marshal(response)
	if err != nil {
		InternalError(w, r, ErrorInternalServer)
	}

	returnError(w, r, http.StatusBadRequest, res)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	response := Response[any]{
		Data:    nil,
		Message: ErrorNotFound,
		Action:  ActionStop,
	}
	res, err := json.Marshal(response)
	if err != nil {
		InternalError(w, r, ErrorInternalServer)
	}

	returnError(w, r, http.StatusNotFound, res)
}

func returnError(w http.ResponseWriter, r *http.Request, code int, response []byte) http.ResponseWriter {
	w = SetContentType(w, r, code)
	_, err := w.Write(response)
	if err != nil {
		InternalError(w, r, ErrorInternalServer)
	}

	return w
}

func InternalError(w http.ResponseWriter, r *http.Request, data interface{}) {
	response, _ := json.Marshal(data)

	w = SetContentType(w, r, http.StatusInternalServerError)
	_, err := w.Write(response)
	if err != nil {
		logger.Error(ErrorInternalServer, err.Error())
	}
}

func SetContentType(w http.ResponseWriter, r *http.Request, code int) http.ResponseWriter {
	w.Header().Set("Content-Type", contentType(r))
	w.WriteHeader(code)
	return w
}

func contentType(r *http.Request) string {
	contentType := r.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/json"
	}
	return contentType
}
