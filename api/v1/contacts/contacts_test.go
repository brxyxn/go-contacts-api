package contacts_test

import (
	"bytes"
	"encoding/json"
	"github.com/brxyxn/go-phonebook-api/internal/constants"
	"github.com/brxyxn/go-phonebook-api/pkg/logger"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/brxyxn/go-phonebook-api/api/v1/contacts"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var handler *contacts.Handler

func init() {
	cfg := logger.LogConfig{
		Out: os.Stdout,
		LogOptions: logger.LogOptions{
			IsDebug:  true,
			LogLevel: "debug",
			CallerOptions: logger.CallerOptions{
				WithCaller:       true,
				WithCustomCaller: true,
				BasePath:         "go-phonebook-api",
			},
			DateTimeOptions: logger.DateTimeOptions{
				WithTime:   true,
				TimeFormat: constants.LongDateTimeFormat,
			},
		},
	}
	simpleLogger := logger.NewLogger(cfg)
	handler = contacts.NewHandler(&simpleLogger)
}

// E:\go-phonebook-api\api\v1\contacts
// E:\go-phonebook-api\api\v1\contacts\api\v1\contacts\contacts.json
func TestGetContacts(t *testing.T) {
	const contactsFile = "./contacts.json"
	var absolutePath, _ = filepath.Abs(contactsFile)
	log.Println(absolutePath)
	req, _ := http.NewRequest("GET", "/v1/contacts", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/contacts", handler.GetContacts).Methods("GET")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetContact(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1/contacts/a2d25c84-8226-4a4f-a9a0-a0a41863f3a8", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/contacts/{id}", handler.GetContact).Methods("GET")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestCreateContact(t *testing.T) {
	contact := contacts.Contact{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "email@example.com",
		Phone:     []string{"123-456-7890"},
	}
	jsonContact, _ := json.Marshal(contact)
	req, _ := http.NewRequest("POST", "/v1/contacts", bytes.NewBuffer(jsonContact))
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/contacts", handler.CreateContact).Methods("POST")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestUpdateContact(t *testing.T) {
	contact := contacts.Contact{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "new@email.com",
		Phone:     []string{"123-456-7890"},
	}
	jsonContact, _ := json.Marshal(contact)
	req, _ := http.NewRequest("PUT", "/v1/contacts/a2d25c84-8226-4a4f-a9a0-a0a41863f3a8", bytes.NewBuffer(jsonContact))
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/contacts/{id}", handler.UpdateContact).Methods("PUT")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteContact(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/v1/contacts/a2d25c8z-8226-4a4z-a9az-a0a4186zf3a8", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/contacts/{id}", handler.DeleteContact).Methods("DELETE")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}
