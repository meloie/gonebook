package models

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"

	"gonebook/internal/ent"
	"gonebook/internal/validation"
)

type ContactPayload struct {
	Name    string `json:"name,omitempty"`
	Phone   string `json:"phone,omitempty"`
	Address string `json:"address,omitempty"`
}

func (c *ContactPayload) Bind(r *http.Request) error {
	if c.Address == "" && c.Phone == "" && c.Name == "" {
		return errors.New("can't create a contect with all empty fields")
	}
	if c.Phone != "" {
		isPhone := validation.ValidatePhoneNumber(c.Phone)
		if isPhone == false {
			return errors.New("phone number must be in +xxxxxxxxxx format")
		}
	}
	return nil
}

type ContactResponsePayload struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func (c *ContactResponsePayload) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewContactsResponseList(contactList []*ent.Contact) []render.Renderer {
	contacts := []render.Renderer{}
	for _, c := range contactList {
		contacts = append(contacts, NewContactsResponse(c))
	}
	return contacts
}

func NewContactsResponse(contact *ent.Contact) render.Renderer {
	return &ContactResponsePayload{
		ID:      contact.ID,
		Name:    contact.Name,
		Phone:   contact.Phone,
		Address: contact.Address,
	}
}
