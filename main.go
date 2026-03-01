package main

import (
	"github.com/CromartyForth/gator/internal/config"
)

//Read the config file.
func main() {
	
	config.Read()
}
//Set the current user to "lane" (actually, you should use your name instead) and update the config file on disk.

//Read the config file again and print the contents of the config struct to the terminal.