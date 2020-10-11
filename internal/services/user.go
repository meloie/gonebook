package services

import (
	"context"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"gonebook/internal/ent"
	"gonebook/internal/ent/user"
	"gonebook/pkg/models"
)

// CreateUser bcyrpt password and save it in databas
func (svc *Service) CreateUser(ctx context.Context, user models.User) (*ent.User, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "failed to bcyrpt password")
	}
	u, err := svc.Database.User.Create().SetUsername(user.Username).SetPassword(string(pass)).Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create user")
	}
	return u, nil
}

func (svc *Service) GetUserByUsername(ctx context.Context, username string) (*ent.User, error) {
	u, err := svc.Database.User.Query().Where(user.UsernameEQ(username)).Only(ctx)
	return u, err
}

func (svc *Service) CheckUserPassword(usr *ent.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(password))
	if err != nil {
		return false
	}
	return true
}
