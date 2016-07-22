package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config contains the Twitter API keys.
type Config struct {
	ConsumerKey,
	ConsumerSecret,
	AccessToken,
	AccessTokenSecret string
}

// New creates a new Config from the given fields in config.json
func New() (*Config, error) {
	var c Config
	file, err := os.Open("./config.json")
	if err != nil {
		return nil, fmt.Errorf("error opening config file")
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&c); err != nil {
		return nil, fmt.Errorf("error decoding config JSON")
	}

	defer file.Close()

	return &c, nil
}
