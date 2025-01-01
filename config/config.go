package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json:"token"`
	BotPrefix string `json:"botPrefix"`
}

func ReadConfig() error {
	fmt.Println("Reading config file...")
	file, err := os.ReadFile("./config.json")
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("Error parsing config file:", err)
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
	return nil
}
