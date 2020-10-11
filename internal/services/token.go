package services

import (
	"context"
	"gonebook/internal/ent"
	"gonebook/pkg/utils"
)

const TokenLength = 16

func (svc *Service) CreateToken(ctx context.Context, user *ent.User) (*ent.Token, error) {
	t, err := svc.Database.Token.Create().SetUserID(user.ID).
		SetValue(utils.RandString(TokenLength)).Save(ctx)
	return t, err
}
