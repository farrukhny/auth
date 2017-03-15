package common

import (
	"encoding/json"
	"log"
	"os"
)

type configurations struct {
	DBHost     string
	DBUser     string
	DBPassword string
	Database   string
}

// APPENV will read env varible to create config file name
var APPENV = os.Getenv("APPENV") + ".json"

// AppConfig will hold configuration values from config file
var AppConfig configurations

// func for Initialize AppConfig
func initConfig() {
	loadConfigFile()
}

// Read config file based on env
func loadConfigFile() {
	file, err := os.Open("config/" + APPENV)
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	AppConfig = configurations{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}
