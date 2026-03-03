package config

import (
	"os"
	"fmt"
	"io"
	"encoding/json"
)


const configFileName = "/.gatorconfig.json"




// Export a Config struct that represents the JSON file structure, including struct tags.
type Config struct {
	DbURL string `json:"db_url"`
	UserName string
}

// Export a Read function that reads the JSON file found at ~/.gatorconfig.json and returns a Config struct. It should read the file from the HOME directory, then decode the JSON string into a new Config struct. I used os.UserHomeDir to get the location of HOME.
func Read() Config {
	fullpath, err:= getConfigFilePath()
	if err != nil {
		fmt.Println("er")
	}

	// open file readonly
	file, err := os.Open(fullpath)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		return Config{}
	}

	// read file ?
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("error reading file")
		return Config{}
	}

	newConfig := Config{}
	err = json.Unmarshal(data, &newConfig)
	if err != nil {
		fmt.Println("error parsing json into struct")
		return Config{}
	}

	return newConfig
}

// Export a SetUser method on the Config struct that writes the config struct to the JSON file after setting the current_user_name field.
func (c *Config) SetUser(userName string) {
	c.UserName = userName
	err := write(*c)
	if err != nil {
		fmt.Printf("error writing file: %v", err)
	}
}


func write(cfg Config) error {
	// get filepath
	fullpath, err:= getConfigFilePath()
	if err != nil {
		fmt.Println("er")
	}

	// convert to json json.MarshalIndent(data, "<prefix>", "<indent>")
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		fmt.Println("error converting to json")
	}

	// write to file
	err = os.WriteFile(fullpath, data, 0644)
	if err != nil {
		fmt.Println("error writing to file")
	}
	return nil
}


func getConfigFilePath() (string, error){
	home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
	fullpath := home + configFileName
	return fullpath, nil
}