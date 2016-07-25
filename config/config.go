package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Open opens the named file for reading.
type Open func(name string) (*os.File, error)

// Config contains the Twitter API keys.
type Config struct {
	ConsumerKey,
	ConsumerSecret,
	AccessToken,
	AccessTokenSecret string
}

// New creates a new Config from the given fields in config.json
func New(o Open) (*Config, error) {
	c := new(Config)
	file, err := o("./config.json")
	if err != nil {
		return nil, fmt.Errorf("error opening config file")
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&c); err != nil {
		return nil, fmt.Errorf("error decoding config JSON")
	}

	defer file.Close()

	return c, nil
}
