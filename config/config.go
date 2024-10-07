package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Subject string `json:"subject"`
}

func GetSubject() string {
	var config Config
	data, err := os.ReadFile("config/config.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatal(err)
	}

	return config.Subject
}
