package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BiteshipAPIKey string
	BiteshipURL	 string
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("error loading .env file")
	}
	return &Config{
		BiteshipAPIKey: os.Getenv("BITESHIP_APIKEY"),
		BiteshipURL : os.Getenv("BITESHIP_URL"),
	}
}