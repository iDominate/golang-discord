package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type ConfigStruct struct {
	Token    string `json:"token"`
	BotToken string `json:"bottoken"`
}

func (c ConfigStruct) String() string {
	return fmt.Sprintf("Token: %s\nBotPrefix: %s", c.Token, c.BotToken)
}

func GetConfig() (*ConfigStruct, error) {
	err := godotenv.Load(".env")

	if err != nil {
		return &ConfigStruct{}, err
	}

	token := os.Getenv("TOKEN")
	botToken := os.Getenv("BOT_PREFIX")

	return &ConfigStruct{Token: token, BotToken: botToken}, nil
}
