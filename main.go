package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

const maxTweetLength int = 280

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "not enough arguments passed to command")
		os.Exit(-1)
	}

	config := flag.String("config", ".config/tweet/config.json", "config file location")
	flag.Parse()

	home := os.Getenv("HOME")
	file, err := ioutil.ReadFile(filepath.Join(home, *config))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening config file: %v", err)
		os.Exit(-1)
	}

	t := twitter{
		client: http.DefaultClient,
	}

	if err := json.Unmarshal(file, &t); err != nil {
		fmt.Fprintf(os.Stderr, "error unmarshalling config: %v", err)
		os.Exit(-1)
	}

	err = t.tweet(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error sending tweet: %v", err)
		os.Exit(-1)
	}

	fmt.Println("Tweet successfully sent!")
}

type twitter struct {
	client *http.Client

	ConsumerKey       string `json:"consumer_key"`
	ConsumerSecret    string `json:"consumer_secret"`
	AccessToken       string `json:"access_token"`
	AccessTokenSecret string `json:"access_token_secret"`
}

// tweet takes Twitter app (https://apps.titter.com) credentials and passes
// them to an OAuth generator along with a given status. It then attempts
// to send a POST request to the Twitter API and handle a response.
func (t twitter) tweet(status string) error {
	if len(status) > maxTweetLength {
		return fmt.Errorf("tweet exceeds %d character limit", maxTweetLength)
	}

	oa := newOAuthDetails(t.ConsumerKey, t.ConsumerSecret, t.AccessToken, t.AccessTokenSecret, status)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s?status=%s", apiURL, encodeStatus(status)), nil)
	if err != nil {
		return fmt.Errorf("error building request: %v", err)
	}
	req.Header.Set("Authorization", oa.String())

	res, err := t.client.Do(req)
	if err != nil {
		return fmt.Errorf("error performing HTTP request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}
	return nil
}
