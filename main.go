package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/danbondd/tweet/config"
	"github.com/danbondd/tweet/tweet"
)

func main() {
	c, err := config.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	t := tweet.NewTweet(http.DefaultClient)
	res, err := t.Send(c, os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error sending tweet: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(res)
}
