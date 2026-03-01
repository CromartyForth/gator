package main

import (
	"fmt"
	"github.com/CromartyForth/gator/internal/config"
)


func main() {
	//Read the config file.
	cnf := config.Read()
	fmt.Println(cnf)

	//Set the current user to "lane" (actually, you should use your name instead) and update the config file on disk.
	cnf.SetUser("Bilbo Baggins")


	//Read the config file again and print the contents of the config struct to the terminal.
	fmt.Println()
}


