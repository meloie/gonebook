package services

import (
	"context"
	"errors"
	"gonebook/internal/ent"
	"gonebook/internal/ent/token"
	"gonebook/pkg/utils"
	"net/http"
	"strings"
)

const (
	TokenLength = 16
)

var UnAuthorizedErr = errors.New("Unauthorized")

func (svc *Service) CreateToken(ctx context.Context, user *ent.User) (*ent.Token, error) {
	t, err := svc.Database.Token.Create().SetUserID(user.ID).
		SetValue(utils.RandString(TokenLength)).Save(ctx)
	return t, err
}

func (svc *Service) GetUserIDByToken(ctx context.Context, tokenValue string) (int, error) {
	t, err := svc.Database.Token.Query().Where(token.ValueEQ(tokenValue)).Only(ctx)
	if err != nil {
		return 0, err
	}
	id, err := t.QueryUser().OnlyID(ctx)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (svc *Service) CurrentUser(ctx context.Context, r http.Request) (int, error) {
	tokenHeader := r.Header.Get("Authorization")
	if tokenHeader == "" {
		return 0, UnAuthorizedErr
	}
	parts := strings.Fields(tokenHeader)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return 0, UnAuthorizedErr
	}
	id, err := svc.GetUserIDByToken(ctx, parts[1])
	if ent.IsNotFound(err) {
		return 0, UnAuthorizedErr
	}
	return id, err
}
