package command

import (
	"fmt"
	"github.com/CromartyForth/gator/internal/config"
)

type State struct {
	Stateptr *config.Config
}

type command struct {
	name string
	arguments []string
}


// maps command.name to handler functions

type Commands struct {
	cmdToHandler map[string]func(*State, command) error
}


// get the cmd.name, matches it to the function call with the cmd.arguments

func (c Commands) runCommmand(s *State, cmd command) error {
	// check ptr is not nil but non empty state struct
	var emptyS = State{}
	if *s == emptyS {
		return fmt.Errorf("state ptr is empty")
	}
	if s.Stateptr.UserName == "" || s.Stateptr.DbURL == "" {
		return fmt.Errorf("config not set")
	}
	// run the command
	c.cmdToHandler[cmd.name](s, cmd)
	return nil
}

func HandlerLogin(s *State, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("no arguments have been provided")
	}

	s.Stateptr.UserName = cmd.arguments[0]
	fmt.Printf("Username set to %v\n", s.Stateptr.UserName)
	return nil
}


// This method registers a new handler function for a command name.

func (c *Commands) Register (name string, f func(*State, command) error) {
	c.cmdToHandler[name] = f
}