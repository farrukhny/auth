package common

import (
	"encoding/json"
	"log"
	"os"
)

type config struct {
	AppName          string `json:"app_name"`
	BaseURL          string `json:"base_url"`
	Port             string `json:"port"`
	Database         string `json:"database"`
	DatabaseHost     string `json:"database_host"`
	DatabaseUser     string `json:"database_user"`
	DatabasePassword string `json:"database_password"`
}

// Read APPENV varible to pass config file name
var APPENV = os.Getenv("APPENV") + ".json"

// AppConfig will hold configuration values from config file
var AppConfig config

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
	AppConfig = config{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}
