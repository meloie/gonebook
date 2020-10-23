package validation

import (
	"context"

	"github.com/rs/zerolog/log"

	"gonebook/internal/ent"
	"gonebook/internal/ent/user"
)

func CheckUserUsernameDuplication(ctx context.Context, db *ent.Client, username string) bool {
	exist, err := db.User.Query().Where(user.UsernameEQ(username)).Exist(ctx)
	if err != nil {
		log.Error().Err(err).Msg("query user failed")
	}
	return exist
}
