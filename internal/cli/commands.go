package cli

import (
	"fmt"
	"strings"

	"github.com/bzelaznicki/gator/internal/config"
)

type state struct {
	Cfg *config.Config
}

type Command struct {
	Name      string
	Arguments []string
}
type Commands struct {
	cmds map[string]func(*state, Command) error
}

func NewCommands() *Commands {
	return &Commands{
		cmds: make(map[string]func(*state, Command) error),
	}
}

func NewState(cfg *config.Config) *state {
	return &state{Cfg: cfg}
}

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

func (c *Commands) Register(name string, f func(*state, Command) error) error {
	name = strings.ToLower(name)
	if _, exists := c.cmds[name]; exists {
		return fmt.Errorf("command %q is already registered", name)
	}

	c.cmds[name] = f
	return nil
}

func (c *Commands) Run(s *state, cmd Command) error {
	cmd.Name = strings.ToLower(cmd.Name)
	handler, exists := c.cmds[cmd.Name]
	if !exists {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}
	return handler(s, cmd)
}
