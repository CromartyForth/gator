package main

// lane doesn't like this!
import _ "github.com/lib/pq"

import (
	"os"
	"fmt"
	"database/sql"
	"github.com/CromartyForth/gator/internal/config"
	"github.com/CromartyForth/gator/internal/command"
	"github.com/CromartyForth/gator/internal/database"
)


func main() {
	// Read the config.json file.
	cnf := config.Read()

	// open a connection to the database
	db, err := sql.Open("postgres", cnf.DbURL)
	if err != nil {
		fmt.Printf("Errror opening database: %v", err)
	}

	dbQueries := database.New(db)
	
	// store *config in state
	newState := command.State{
		Db: dbQueries,
		Stateptr: &cnf,
	}


	//Creates a new instance of the commands struct with an initialized map of handler functions.
	newCommands := command.Commands{}
	newCommands.CmdToHandler = make(map[string]func(*command.State, command.Command) error)

	// Registers a handler function for commands.
	newCommands.Register("login", command.HandlerLogin)
	newCommands.Register("register", command.HandlerRegister)
	newCommands.Register("reset", command.HandlerReset)
	newCommands.Register("users", command.HandlerUsers)


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
	
	err = newCommands.RunCommand(&newState, userCmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%+v", newState.Stateptr)

}


