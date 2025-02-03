package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/bzelaznicki/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerRegister(s *state, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("registration command requires a username. Usage: register <username>")
	}

	username := cmd.Arguments[0]

	_, err := s.db.GetUser(context.Background(), username)
	if err == nil {

		return fmt.Errorf("user %s already exists", username)

	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Arguments[0],
	}

	userInfo, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}
	err = s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("User %s successfully created!\n", username)
	fmt.Printf("ID: %s\n Name: %s\n CreatedAt: %s\n UpdatedAt: %s\n", userInfo.ID, userInfo.Name, userInfo.CreatedAt, userInfo.UpdatedAt)
	return nil
}
