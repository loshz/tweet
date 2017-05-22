package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "not enough arguments passed to command")
		os.Exit(-1)
	}

	homeDir := os.Getenv("HOME")
	config, err := NewConfig(homeDir, configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}

	tweet := NewTweet(http.DefaultClient, http.NewRequest)
	err = tweet.Send(config, os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error sending tweet: %v\n", err)
		os.Exit(-1)
	}

	fmt.Println("Tweet successfully sent!")
}
