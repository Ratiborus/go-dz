package config

import (
	"fmt"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key    string
	apiUrl url.URL
}

func NewConfig() *Config {
	godotenv.Load(".env")
	key := os.Getenv("KEY")
	if key == "" {
		panic("Encryption Key(KEY) is not found in env file")
	}
	link := os.Getenv("API_URL")
	if link == "" {
		panic("Api url(API_URL) is not found in env file")
	}
	u, err := url.Parse(link)
	if err != nil {
		panic(fmt.Sprintf("Can not parse api url(API_URL) to url string %v", err))
	}
	return &Config{Key: key, apiUrl: *u}
}

func (c *Config) BinUrl() *url.URL {
	return c.apiUrl.JoinPath("b")
}
