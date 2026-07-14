package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	godotenv.Load(".env")
	key := os.Getenv("KEY")
	if key == "" {
		panic("Encryption Key(KEY) is not found in env file")
	}
	return &Config{Key: key}
}
