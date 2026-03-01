package config

import (
	"os"
	"fmt"
)


const configFileName = ".gatorconfig.json"

// Export a Config struct that represents the JSON file structure, including struct tags.
type Config struct {
	UserName string `json:"userName"`
	DbURL string `json:"db_url"`
}

// Export a Read function that reads the JSON file found at ~/.gatorconfig.json and returns a Config struct. It should read the file from the HOME directory, then decode the JSON string into a new Config struct. I used os.UserHomeDir to get the location of HOME.
func Read() Config {
	home, err := os.UserHomeDir()
		if err != nil {
			return Config{}
		}
	fmt.Println(home)
	
	return Config{}
}

// Export a SetUser method on the Config struct that writes the config struct to the JSON file after setting the current_user_name field.
func (c *Config) SetUser(userName string) {

}

func getConfigFilePath() (string, error){
	return "", nil
}

func write(cfg Config) error {
	return nil
}

