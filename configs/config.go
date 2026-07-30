package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host     string
	Password string
	File     string
	SendConfig
}

type SendConfig struct {
	DefaultEmail  string
	SendToken     string
	ReserverToken string
}

func GetConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	return &Config{
		Host:     os.Getenv("HOST"),
		Password: os.Getenv("PASSWORD"),
		File:     os.Getenv("FILE"),
		SendConfig: SendConfig{
			DefaultEmail:  os.Getenv("DEFAULT_ADDRESS"),
			SendToken:     os.Getenv("SEND_TOKEN"),
			ReserverToken: os.Getenv("RESERVE_TOKEN"),
		},
	}, nil
}
