package webserver

import (
	"errors"
	"gonebook/internal/ent"
	"gonebook/pkg/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

func (srv *WebServer) ContactsList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := srv.service.CurrentUser(ctx, r)
	if err != nil {
		render.Render(w, r, models.ErrUnathorized(err))
		return
	}
	userContacts, err := srv.service.GetUserContacts(ctx, userID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get contacts")
		internalError(w)
		return
	}
	contacts := models.NewContactsResponseList(userContacts)
	render.RenderList(w, r, contacts)
}

func (srv *WebServer) CreateContact(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := srv.service.CurrentUser(ctx, r)
	if err != nil {
		render.Render(w, r, models.ErrUnathorized(err))
		return
	}
	var data models.ContactPayload
	err = render.Bind(r, &data)
	if err != nil {
		render.Render(w, r, models.ErrInvalidRequest(err))
		return
	}
	contact, err := srv.service.CreateContact(ctx, &data, userID)
	if err != nil {
		log.Error().Err(err).Msg("failed to create contact")
		internalError(w)
		return
	}
	render.Status(r, 201)
	render.Render(w, r, models.NewContactsResponse(contact))
}

func (srv *WebServer) ContactDetails(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := srv.service.CurrentUser(ctx, r)
	if err != nil {
		render.Render(w, r, models.ErrUnathorized(err))
		return
	}
	contactID, err := strconv.Atoi(chi.URLParam(r, "contactId"))
	if err != nil {
		render.Render(w, r, models.ErrNotFound(errors.New("contact not found")))
		return
	}
	contact, err := srv.service.GetContact(ctx, contactID)
	if ent.IsNotFound(err) {
		render.Render(w, r, models.ErrNotFound(errors.New("contact not found")))
		return
	}
	if contact.Edges.Owner.ID != userID {
		render.Render(w, r, models.ErrAccessDenied(errors.New("you are not owner of this contact")))
		return
	}
	render.Render(w, r, models.NewContactsResponse(contact))
}

func (srv *WebServer) UpdateContact(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := srv.service.CurrentUser(ctx, r)
	if err != nil {
		render.Render(w, r, models.ErrUnathorized(err))
		return
	}
	contactID, err := strconv.Atoi(chi.URLParam(r, "contactId"))
	if err != nil {
		render.Render(w, r, models.ErrNotFound(errors.New("contact not found")))
		return
	}
	contact, err := srv.service.GetContact(ctx, contactID)
	if ent.IsNotFound(err) {
		render.Render(w, r, models.ErrNotFound(errors.New("contact not found")))
		return
	}

	if contact.Edges.Owner.ID != userID {
		render.Render(w, r, models.ErrAccessDenied(errors.New("you are not owner of this contact")))
		return
	}

	var data models.ContactPayload
	err = render.Bind(r, &data)
	if err != nil {
		render.Render(w, r, models.ErrInvalidRequest(err))
		return
	}

	updatedContact, err := srv.service.UpdateContact(ctx, contact, data)
	if err != nil {
		log.Error().Err(err).Int("contact_id", contactID).Msg("failed to update contact")
		internalError(w)
		return
	}
	render.Render(w, r, models.NewContactsResponse(updatedContact))
}

func (srv *WebServer) DeleteContact(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := srv.service.CurrentUser(ctx, r)
	if err != nil {
		render.Render(w, r, models.ErrUnathorized(err))
		return
	}
	contactID, err := strconv.Atoi(chi.URLParam(r, "contactId"))
	if err != nil {
		log.Info().Str("contact_id", chi.URLParam(r, "contactId")).Msg("contactId is invalid")
		render.Render(w, r, models.ErrNotFound(errors.New("contact not found")))
		return
	}
	contact, err := srv.service.GetContact(ctx, contactID)
	if ent.IsNotFound(err) {
		log.Info().Err(err).Msg("wrong contactId")
		render.Render(w, r, models.ErrNotFound(errors.New("contact not found")))
		return
	}
	if contact.Edges.Owner.ID != userID {
		render.Render(w, r, models.ErrAccessDenied(errors.New("you are not owner of this contact")))
		return
	}

	err = srv.service.DeleteContact(ctx, contact)
	if err != nil {
		log.Error().Err(err).Int("contact_id", contactID).Msg("failed to update contact")
		internalError(w)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
