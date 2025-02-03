package cli

import (
	"context"
	"fmt"
)

func HandlerLogin(s *state, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("login command requires a username. Usage: login <username>")
	}
	username := cmd.Arguments[0]

	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {

		return fmt.Errorf("user %s not found", username)

	}
	err = s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully logged in as %s\n", s.cfg.CurrentUserName)
	return nil
}
