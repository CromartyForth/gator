package main

import (
	"os"
	"fmt"
	"github.com/CromartyForth/gator/internal/config"
	"github.com/CromartyForth/gator/internal/command"
)


func main() {
	// Read the config file.
	cnf := config.Read()

	// store *config in state
	newState := command.State{
		Stateptr: &cnf,
	}
	
	//Creates a new instance of the commands struct with an initialized map of handler functions.
	newCommands := command.Commands{}
	newCommands.CmdToHandler = make(map[string]func(*command.State, command.Command) error)

	// Registers a handler function for the login command.
	newCommands.Register("login", command.HandlerLogin)

	// get cmdln args
	if len(os.Args) < 2 {
		fmt.Println("not enough arguments supplied")
		os.Exit(1)
	}
	// split args into slice (field)
	userCmd := command.Command {
		Name: os.Args[1],
		Arguments: os.Args[2:],
	}
	
	err := newCommands.RunCommand(&newState, userCmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(newState.Stateptr)

}


