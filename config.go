package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

const configFile string = "/.config/tweet/config.json"

// Config contains the Twitter API keys.
type Config struct {
	ConsumerKey,
	ConsumerSecret,
	AccessToken,
	AccessTokenSecret string
}

// NewConfig creates a new Config from the given fields in config.json
func NewConfig(homeDir, filename string) (*Config, error) {
	c := new(Config)

	file, err := ioutil.ReadFile(filepath.Join(homeDir, filename))
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %v", err)
	}

	if err := json.Unmarshal(file, &c); err != nil {
		return nil, fmt.Errorf("error unmarshalling config JSON: %v", err)
	}

	return c, nil
}
