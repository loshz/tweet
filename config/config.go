package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// openFile opens the named file for reading.
type openFile func(name string) (io.ReadCloser, error)

// FileReader l
func FileReader(name string) (io.ReadCloser, error) {
	return os.Open(name)
}

// Config contains the Twitter API keys.
type Config struct {
	ConsumerKey,
	ConsumerSecret,
	AccessToken,
	AccessTokenSecret string
}

// Decoder l
type Decoder interface {
	Decode(v interface{}) error
}

type decoderFactory func(r io.Reader) Decoder

// JSONDecoderFactory l
func JSONDecoderFactory(r io.Reader) Decoder {
	return json.NewDecoder(r)
}

// New creates a new Config from the given fields in config.json
func New(open openFile, d decoderFactory) (*Config, error) {
	c := new(Config)
	file, err := open("./config.json")
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
