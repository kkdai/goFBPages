package main

import (
	"encoding/json"
	"log"
	"os"
)

// It is the config for facebook album communication.
type CommunicationConfig struct {
	//Token for facebook connection. The lifecycle of token is shortly, please remember to refresh it on  https://developers.facebook.com/tools/explorer?method=GET
	Token string
}

func LoadConfig() CommunicationConfig {
	file, _ := os.Open("conf.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := CommunicationConfig{}
	err := decoder.Decode(&configuration)
	if err != nil || len(configuration.Token) == 0 {
		log.Fatalln(" No config file or no token, please provide conf.json in the same folder.")
	}
	return configuration
}
