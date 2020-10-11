package validation

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"

	"gonebook/internal/ent"
	"gonebook/internal/ent/user"
	"gonebook/pkg/models"
)

func CheckUserUsernameDuplication(ctx context.Context, db *ent.Client, username string) bool {
	exist, err := db.User.Query().Where(user.UsernameEQ(username)).Exist(ctx)
	if err != nil {
		log.Error().Err(err).Msg("query user failed")
	}
	return exist
}

func IsRegistrationInfoValid(usr models.User) error {
	if usr.Username == "" {
		return errors.New("username can't be empty")
	}
	if usr.Password == "" {
		return errors.New("password can't be empty")
	}
	return nil
}
