package common

import (
	"encoding/json"
	"log"
	"os"
)

type config struct {
	AppName          string `json:"app_name,omitempty"`
	BaseURL          string `json:"base_url,omitempty"`
	Port             string `json:"port,omitempty"`
	Database         string `json:"database,omitempty"`
	DatabaseHost     string `json:"database_host,omitempty"`
	DatabaseUser     string `json:"database_user,omitempty"`
	DatabasePassword string `json:"database_password,omitempty"`
}

// Read APPENV varible to pass config file name

var APPENV string = os.Getenv("APPENV")

// AppConfig will hold configuration values from config file
var AppConfig config

// func for Initialize AppConfig
func initConfig() {
	loadConfigFile()
}

// Read config file based on env
func loadConfigFile() {
	if APPENV == "" {
		APPENV = "dev"
	}
	file, err := os.Open("config/" + APPENV + ".json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	AppConfig = config{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}
