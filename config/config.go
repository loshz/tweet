package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/danbondd/tweet/utils"
)

const configFile string = "/.config/tweet/config.json"

// Config contains the Twitter API keys.
type Config struct {
	ConsumerKey,
	ConsumerSecret,
	AccessToken,
	AccessTokenSecret string
}

// New creates a new Config from the given fields in config.json
func New(open utils.OpenFile, d utils.JSONDecoder) (*Config, error) {
	c := new(Config)

	homeDir := os.Getenv("HOME")
	if len(homeDir) == 0 {
		return nil, errors.New("home directory not set")
	}

	path := []string{homeDir, configFile}
	file, err := open(strings.Join(path, ""))
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %v", err)
	}
	defer file.Close()

	decoder := d(file)
	if err := decoder.Decode(&c); err != nil {
		return nil, fmt.Errorf("error decoding config JSON: %v", err)
	}

	return c, nil
}
