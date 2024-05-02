package contacts

import (
	"encoding/json"
	"github.com/brxyxn/go-phonebook-api/internal/response"
	"github.com/gorilla/mux"
	"net/http"
)

const contactsFile = "./api/v1/contacts/contacts.json"

// const contactsFile = "./contacts.json" // uncomment this line when running tests

const (
	SuccessContacts = "contacts retrieved successfully"
	SuccessContact  = "contact retrieved successfully"
	SuccessCreate   = "contact created successfully"
	SuccessUpdate   = "contact updated successfully"
	SuccessDelete   = "contact deleted successfully"
)

type Contact struct {
	ID        string   `json:"id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Email     string   `json:"email"`
	Phone     []string `json:"phone"`
}

type Contacts []Contact

type Response struct {
	Contacts interface{} `json:"contacts"`
	Total    int         `json:"total"`
}

/*
swagger:route GET /v1/contacts contacts GetContacts
Return a list of all contacts
responses:

	200
	404
	500

...
GetContacts returns all contacts
e.g. [GET] /v1/contacts
*/
func (c *Handler) GetContacts(w http.ResponseWriter, r *http.Request) {
	var res Response

	data, err := getAllContacts()
	if err != nil {
		c.logger.Error().Err(err).Msg("error getting contacts")
		response.InternalError(w, r, data)
		return
	}

	if len(data) < 1 {
		c.logger.Error().Int("total", len(data)).Msg("no contacts found")
		response.NotFound(w, r)
		return
	}

	res.Contacts = data
	res.Total = len(data)

	response.Success(w, r, res, SuccessContacts)
	c.logger.Info().Int("total", res.Total).Msg(SuccessContacts)
}

/*
swagger:route GET /v1/contacts/{id} contacts GetContact
Return the details of a single contact
responses:

	200
	404
	500

...
GetContact returns a single contact by ID path parameter
e.g. [GET] /v1/contacts/c32b6462-7b0b-462a-b7af-21b1ae3f273b
*/
func (c *Handler) GetContact(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var err error

	data, _, err := getContactByID(id)
	if err != nil {
		c.logger.Error().Str("id", id).Err(err).Msg("error getting contact")
		response.NotFound(w, r)
		return
	}

	response.Success(w, r, data, SuccessContact)
	c.logger.Info().Str("id", id).Msg(SuccessContact)
}

/*
swagger:route POST /v1/contacts/{id} contacts CreateContact
Return the created contact including the details of it if the request is successful
responses:

	200
	400
	500

...
CreateContact creates a new contact
e.g. [POST] /v1/contacts
request body: { "first_name": "John", "last_name": "Doe", "email": "email@example.com", "phone": ["123-456-7890"] }
*/
func (c *Handler) CreateContact(w http.ResponseWriter, r *http.Request) {
	var contact Contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	c.logger.Debug().Any("contact", contact).Msg("creating contact")
	if err != nil {
		c.logger.Error().Any("contact", contact).Err(err).Msg("error decoding contact")
		response.Error(w, r, err.Error())
		return
	}

	newContact, err := createContact(contact)
	if err != nil {
		c.logger.Error().Any("contact", newContact).Err(err).Msg("failed to create contact")
		response.Error(w, r, err.Error())
		return
	}

	response.Created(w, r, newContact, SuccessCreate)
	c.logger.Info().Any("newContact", newContact).Msg(SuccessCreate)
}

/*
swagger:route PUT /v1/contacts/{id} contacts CreateContact
Return the updated contact including the details of it if the request is successful
responses:

	200
	400
	500

...
UpdateContact updates a contact by ID path parameter and request body
e.g. [PUT] /v1/contacts/c32b6462-7b0b-462a-b7af-21b1ae3f273b
request body: { "first_name": "John", "last_name": "Doe", "email": "new@email.com", "phone": ["123-456-7890"] }
*/
func (c *Handler) UpdateContact(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var updated Contact
	err := json.NewDecoder(r.Body).Decode(&updated)
	c.logger.Info().Any("updated", updated).Msg("updating contact")
	if err != nil {
		c.logger.Error().Msg(err.Error())
		response.Error(w, r, err.Error())
		return
	}

	updatedContact, err := updateContact(id, updated)
	if err != nil {
		c.logger.Error().Msg(err.Error())
		response.Error(w, r, err.Error())
		return
	}

	response.Success(w, r, updatedContact, SuccessUpdate)
	c.logger.Info().Any("updatedContact", updatedContact).Msg(SuccessUpdate)
}

/*
swagger:route GET /v1/contacts/{id} contacts GetContact
Return the details of a single contact
responses:

	204
	404
	500

...
DeleteContact deletes a contact by ID path parameter
e.g. [DELETE] /v1/contacts/c32b6462-7b0b-462a-b7af-21b1ae3f273b
*/
func (c *Handler) DeleteContact(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	c.logger.Info().Str("id", id).Msg("deleting contact")

	err := deleteContact(id)
	if err != nil {
		c.logger.Error().Str("id", id).Err(err).Msg("failed to delete contact")
		if err.Error() == "contact not found" {
			response.NotFound(w, r)
			return
		}
		response.Error(w, r, err.Error())
		return
	}

	response.Deleted(w)
	c.logger.Info().Msg(SuccessDelete)
}
