package webserver

import (
	"gonebook/pkg/models"
	"net/http"

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
	render.Render(w, r, models.NewContactsResponse(contact))
}
