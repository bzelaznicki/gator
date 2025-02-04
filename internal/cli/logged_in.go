package cli

import (
	"context"

	"github.com/bzelaznicki/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd Command, user database.User) error) func(*state, Command) error {

	return func(s *state, cmd Command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)

	}
}
