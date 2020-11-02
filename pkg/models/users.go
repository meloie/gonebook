package models

import (
	"errors"
	"github.com/meloie/gonebook/internal/ent"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Bind on UserPayload will run after the unmarshalling is complete, its
// a good time to focus some post-processing after a decoding.
func (u *User) Bind(r *http.Request) error {
	if u.Username == "" {
		return errors.New("username can't be empty")
	}
	if u.Password == "" {
		return errors.New("password can't be empty")
	}
	return nil
}

type UserPayload struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func NewUserPayloadResponse(user *ent.User) *UserPayload {
	return &UserPayload{
		ID:       user.ID,
		Username: user.Username,
	}
}

func (u *UserPayload) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type TokenPayloadResponse struct {
	Token string `json:"token"`
}

func (t *TokenPayloadResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
