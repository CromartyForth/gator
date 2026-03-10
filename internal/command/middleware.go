package command

import (
	"os"
	"fmt"
	"context"
	"github.com/CromartyForth/gator/internal/database"
)

func MiddlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {

	return func(s *State, cmd Command) error {
		contextBackground := context.Background()
		getuser, err := s.Db.GetUser(contextBackground, s.Stateptr.UserName)
		if err != nil {
			fmt.Printf("User %v does not exist.", s.Stateptr.UserName)
			os.Exit(1)
		}
		return handler(s, cmd, getuser)
	}
	
} 

/*
func addPrefix(prefix string) func(string) string {
	return func(name string) string {
		return prefix + " " + name
	}
}
*/
