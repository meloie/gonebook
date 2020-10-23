package webserver

import (
	"errors"
	"fmt"
	"gonebook/internal/validation"
	"gonebook/pkg/models"
	"net/http"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

func internalError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - Error!"))
}
func (srv *WebServer) RegisterUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var data models.User
	err := render.Bind(r, &data)
	if err != nil {
		render.Render(w, r, models.ErrInvalidRequest(err))
		return
	}
	exists := validation.CheckUserUsernameDuplication(ctx, srv.service.Database, data.Username)
	if exists == true {
		render.Render(w, r, models.ErrInvalidRequest(fmt.Errorf("username \"%s\" already has been taken", data.Username)))
		return
	}
	u, err := srv.service.CreateUser(ctx, data)
	if err != nil {
		log.Error().Err(err).Msg("failed to create user")
		internalError(w)
		return
	}
	render.Render(w, r, models.NewUserPayloadResponse(u))
}

func (srv *WebServer) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var data models.User
	err := render.Bind(r, &data)
	if err != nil {
		log.Error().Err(err).Msg("failed to decode request body")
		render.Render(w, r, models.ErrInvalidRequest(err))
		return
	}
	u, err := srv.service.GetUserByUsername(ctx, data.Username)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user by username")
		internalError(w)
		return
	}
	v := srv.service.CheckUserPassword(u, data.Password)
	if v == true {
		token, err := srv.service.CreateToken(ctx, u)
		if err != nil {
			log.Error().Err(err).Msg("failed to create token")
		}
		render.Render(w, r, &models.TokenPayloadResponse{token.Value})
		return
	}
	render.Render(w, r, models.ErrUnathorized(errors.New("Username and/or password is wrong!")))

}
