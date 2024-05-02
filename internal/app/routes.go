package app

import (
	s "github.com/brxyxn/go-phonebook-api/api/status"
	c "github.com/brxyxn/go-phonebook-api/api/v1/contacts"
)

func (a *Config) setupRoutes() {
	status := s.NewHandlers()
	a.Router.HandleFunc("/status", status.GetStatus).Methods("GET")

	// Contacts
	contacts := c.NewHandler(a.Logger)
	a.Router.HandleFunc("/v1/contacts", contacts.GetContacts).Methods("GET")           // GetAll
	a.Router.HandleFunc("/v1/contacts/{id}", contacts.GetContact).Methods("GET")       // GetById
	a.Router.HandleFunc("/v1/contacts", contacts.CreateContact).Methods("POST")        // Create
	a.Router.HandleFunc("/v1/contacts/{id}", contacts.UpdateContact).Methods("PUT")    // Update
	a.Router.HandleFunc("/v1/contacts/{id}", contacts.DeleteContact).Methods("DELETE") // Delete
}
