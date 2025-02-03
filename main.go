package main

import (
	"fmt"
	"os"

	"github.com/bzelaznicki/gator/internal/cli"
	"github.com/bzelaznicki/gator/internal/config"
)

func main() {
	// Load config
	cfg, err := config.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read config: %v\n", err)
		os.Exit(1)
	}

	// Initialize state
	state := cli.NewState(&cfg)

	// Create commands instance using `newCommands`
	cmds := cli.NewCommands() // Adjust capitalization depending on visibility.

	// Register the login handler
	err = cmds.Register("login", cli.HandlerLogin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to register the login command: %v\n", err)
		os.Exit(1)
	}

	// Ensure enough arguments are provided
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "not enough arguments provided")
		os.Exit(1)
	}

	// Parse command line arguments
	name := os.Args[1]
	args := os.Args[2:]

	// Create the command
	cmd := cli.Command{
		Name:      name,
		Arguments: args,
	}

	// Run the command
	err = cmds.Run(state, cmd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	// Exit successfully
	os.Exit(0)
}
