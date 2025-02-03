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
	// Initialize commands map
	commands := &Commands{
		cmds: make(map[string]func(*state, Command) error),
	}

	// Register command handlers
	commands.Register("login", HandlerLogin)
	//commands.Register("register", HandlerRegister)

	return commands
}

func (c *Commands) Register(name string, f func(*state, Command) error) error {
	name = strings.ToLower(name)
	if _, exists := c.cmds[name]; exists {
		return fmt.Errorf("command %q is already registered", name)
	}

	c.cmds[name] = f
	return nil
}

func NewState(cfg *config.Config) *state {
	return &state{Cfg: cfg}
}

func (c *Commands) Run(s *state, cmd Command) error {
	cmd.Name = strings.ToLower(cmd.Name)
	handler, exists := c.cmds[cmd.Name]
	if !exists {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}
	return handler(s, cmd)
}
