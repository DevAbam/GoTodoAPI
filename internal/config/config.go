package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
}

func LoadEnv() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env not found, using environment variables")
	}

	cfg := Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
	}

	return &cfg, nil
}
