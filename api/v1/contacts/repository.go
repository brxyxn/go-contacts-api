package contacts

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"os"
	"path/filepath"
)

var absolutePath, _ = filepath.Abs(contactsFile) // temporary solution, this should be handled with persisted data storage

func getAllContacts() (Contacts, error) {
	var fileData []byte
	var data Contacts
	var err error
	fileData, err = os.ReadFile(absolutePath)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(fileData, &data)
	return data, err
}

func getContactByID(id string) (*Contact, int, error) {
	contacts, err := getAllContacts()
	if err != nil {
		return nil, -1, err
	}

	for i, p := range contacts {
		if p.ID == id {
			return &p, i, nil
		}
	}
	return nil, -1, fmt.Errorf("contact not found")
}

func createContact(newContact Contact) (Contact, error) {
	// Read the existing contacts
	existingContacts, err := getAllContacts()
	if err != nil {
		return newContact, err
	}

	err = validateContact(newContact, true)
	if err != nil {
		return newContact, err
	}

	newContact.ID = generateID()

	// Append the new contact
	existingContacts = append(existingContacts, newContact)
	return newContact, writeContacts(existingContacts, contactsFile)
}

func updateContact(id string, updatedContact Contact) (*Contact, error) {
	contacts, err := getAllContacts()
	if err != nil {
		return nil, err
	}

	_, index, err := getContactByID(id)
	contact := contacts[index]
	if err != nil {
		return nil, err
	}

	err = validateContact(updatedContact, false)
	if err != nil {
		return nil, err
	}

	if updatedContact.ID != "" {
		contact.ID = updatedContact.ID
	}

	if updatedContact.FirstName != "" {
		contact.FirstName = updatedContact.FirstName
	}

	if updatedContact.LastName != "" {
		contact.LastName = updatedContact.LastName
	}

	if updatedContact.Email == "" {
		return nil, fmt.Errorf("email is required")
	}

	if len(updatedContact.Phone) == 0 {
		return nil, fmt.Errorf("phone is required")
	}

	contact.Email = updatedContact.Email
	contact.Phone = updatedContact.Phone

	contacts[index] = contact

	err = writeContacts(contacts, contactsFile)
	if err != nil {
		return nil, err
	}

	newContact := contacts[index]
	return &newContact, nil
}

func deleteContact(id string) error {
	contacts, err := getAllContacts()
	if err != nil {
		return err
	}

	_, index, err := getContactByID(id)
	if err != nil {
		return err
	}

	contacts = append(contacts[:index], contacts[index+1:]...)

	err = writeContacts(contacts, contactsFile)
	if err != nil {
		return err
	}

	return nil
}

func writeContacts(contacts Contacts, filename string) error {
	// Append the new contact
	contacts = append(contacts)

	// Marshal the updated contacts
	data, err := json.MarshalIndent(contacts, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func validateContact(contact Contact, isNew bool) error {
	if contact.ID != "" {
		return fmt.Errorf("id is auto-generated")
	}

	if contact.Email == "" {
		return fmt.Errorf("email is required")
	}

	if len(contact.Phone) == 0 {
		return fmt.Errorf("phone is required")
	}

	// If the contact is new, the first and last name are required
	if !isNew {
		return nil
	}

	if contact.FirstName == "" {
		return fmt.Errorf("first_name is required")
	}

	if contact.LastName == "" {
		return fmt.Errorf("last_name is required")
	}

	return nil
}

// generate UUID
func generateID() string {
	return uuid.NewString()
}
