package cli

import (
	"fmt"
)

func HandlerLogin(s *state, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("login command requires a username. Usage: login <username>")
	}

	err := s.Cfg.SetUser(cmd.Arguments[0])
	if err != nil {
		return err
	}

	fmt.Printf("Successfully logged in as %s\n", s.Cfg.CurrentUserName)
	return nil
}
