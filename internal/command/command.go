package command

import (
	
	"fmt"
	"github.com/CromartyForth/gator/internal/config"
	"github.com/CromartyForth/gator/internal/database"
)

type State struct {
	Db  *database.Queries
	Stateptr *config.Config
}

type Command struct {
	Name string
	Arguments []string
}


// maps command.name to handler functions
type Commands struct {
	CmdToHandler map[string]func(*State, Command) error
}



// get the cmd.name, matches it to the function call with the cmd.arguments
func (c Commands) RunCommand(s *State, cmd Command) error {
	// check ptr is not nil but non empty state struct
	var emptyS = State{}
	if *s == emptyS {
		return fmt.Errorf("state ptr is empty")
	}
	if s.Stateptr.UserName == "" || s.Stateptr.DbURL == "" {
		return fmt.Errorf("config not set")
	}
	// run the command
	err := c.CmdToHandler[cmd.Name](s, cmd)
	if err != nil {
		return err
	}
	return nil
}

// This method registers a new handler function for a command name.
func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.CmdToHandler[name] = f
}

























