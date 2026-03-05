package command

import (
	"os"
	"fmt"
	"github.com/google/uuid"
	"context"
	"time"
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

func HandlerLogin(s *State, cmd Command) error {
	// ensure username in args
	if len(cmd.Arguments) < 1 {
		return fmt.Errorf("username is required")
	}

	// ensure username is in database
	contextBackground := context.Background()
	user, err := s.Db.GetUser(contextBackground, cmd.Arguments[0])
	if err != nil {
		fmt.Printf("User %v does not exist.", cmd.Arguments[0])
		os.Exit(1)
	}

	s.Stateptr.SetUser(user.Name)

	//s.Stateptr.UserName = cmd.Arguments[0]
	fmt.Printf("Username set to %v\n", s.Stateptr.UserName)
	return nil
}


func HandlerRegister(s *State, cmd Command) error {
	// ensure username in args
	if len(cmd.Arguments) < 1 {
		return fmt.Errorf("username is required")
	}

	// Create a new user in the database. It should have access to the CreateUser query through the state -> db struct.
	contextBackground := context.Background()
	userArgs := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.Arguments[0],
	}
	// Create a new user in the database. It should have access to the CreateUser query through the state -> db struct.
	newUser, err := s.Db.CreateUser(contextBackground, userArgs)
	if err != nil {
		fmt.Printf("User with that name already exists: %v", err)
		os.Exit(1)
	}
	
	// Set the current user in the config to the given name.
	s.Stateptr.SetUser(cmd.Arguments[0])

	// Print a message that the user was created
	fmt.Printf("Username set to %v\n", s.Stateptr.UserName)
	fmt.Printf("%+v", newUser)
	return nil
}

func HandlerReset(s *State, cmd Command) error {
	contextBackground := context.Background()
	if err := s.Db.DeleteAllUsers(contextBackground); err != nil {
		return fmt.Errorf("error deleting user table: %v", err)
	}
	
	return nil
}

func HandlerUsers(s *State, cmd Command) error {
	contextBackground := context.Background()

	// get all the users
	users, err := s.Db.GetUsers(contextBackground)
	if err != nil {
		return fmt.Errorf("error getting all users: %v", err)
	}

	// Whois the current user.
	currentUser := s.Stateptr.UserName
	
	// print out users
	for _, user := range users {
		if user == currentUser {
			fmt.Printf("%v (current)\n", user)
		} else {
			fmt.Println(user)
		}
	}

	return nil
}

func HandlerAgg(s *State, cmd Command) error {
	// Add an agg command. Later this will be our long-running aggregator service. For now, we'll just use it to fetch a single feed and ensure our parsing works. It should fetch the feed found at https://www.wagslane.dev/index.xml and print the entire struct to the console.
	contextBackground := context.Background()
	fetchedFeed, err := fetchFeed(contextBackground, "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Printf("%+v", fetchedFeed)

	return nil
}
