package cli

import (
	"testing"

	"github.com/bzelaznicki/gator/internal/config"
)

func TestHandlerLogin(t *testing.T) {
	cfg := &config.Config{}
	s := NewState(cfg)
	cmd := Command{Name: "login", Arguments: []string{"testuser"}}

	err := HandlerLogin(s, cmd)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if s.Cfg.CurrentUserName != "testuser" {
		t.Fatalf("expected username to be 'testuser', got %s", s.Cfg.CurrentUserName)
	}
}

func TestHandlerLoginNoUsername(t *testing.T) {
	cfg := &config.Config{}
	s := &state{Cfg: cfg}
	cmd := Command{Name: "login", Arguments: []string{}}

	err := HandlerLogin(s, cmd)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	expectedError := "provide a username"
	if err.Error() != expectedError {
		t.Fatalf("expected error %q, got %q", expectedError, err.Error())
	}
}

func TestCommandsRegisterAndRun(t *testing.T) {
	cmds := NewCommands()
	err := cmds.Register("login", HandlerLogin)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	cfg := &config.Config{}
	s := &state{Cfg: cfg}
	cmd := Command{Name: "login", Arguments: []string{"testuser"}}

	err = cmds.Run(s, cmd)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if s.Cfg.CurrentUserName != "testuser" {
		t.Fatalf("expected username to be 'testuser', got %s", s.Cfg.CurrentUserName)
	}
}

func TestCommandsRegisterDuplicate(t *testing.T) {
	cmds := NewCommands()
	err := cmds.Register("login", HandlerLogin)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = cmds.Register("login", HandlerLogin)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	expectedError := `command "login" is already registered`
	if err.Error() != expectedError {
		t.Fatalf("expected error %q, got %q", expectedError, err.Error())
	}
}

func TestCommandsRunUnknownCommand(t *testing.T) {
	cmds := NewCommands()
	cfg := &config.Config{}
	s := &state{Cfg: cfg}
	cmd := Command{Name: "unknown", Arguments: []string{}}

	err := cmds.Run(s, cmd)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	expectedError := "unknown command: unknown"
	if err.Error() != expectedError {
		t.Fatalf("expected error %q, got %q", expectedError, err.Error())
	}
}
