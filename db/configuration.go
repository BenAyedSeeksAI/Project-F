package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	ScriptCreateFile string `json:"scriptCreateFile"`
	ScriptSeedFile   string `json:"scriptSeedFile"`
	Host             string `json:"host"`
	Port             int    `json:"port"`
	User             string `json:"user"`
	Password         string `json:"password"`
	DBName           string `json:"dbname"`
}

func GetDBConfiguration(jsonFilePath string) *Config {
	// Read the entire file
	jsonData, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// Create a Config struct to store the values
	var config Config
	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		log.Fatal(err)
	}
	return &config
}
