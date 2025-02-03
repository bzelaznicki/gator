package cli

import (
	"context"
	"fmt"
)

func HandlerReset(s *state, cmd Command) error {
	if len(cmd.Arguments) > 0 {
		return fmt.Errorf("this function does not take arguments")

	}

	err := s.db.ResetDatabaseState(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Database was reset successfully!")
	return nil
}
