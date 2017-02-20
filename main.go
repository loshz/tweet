package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/danbondd/tweet/config"
	"github.com/danbondd/tweet/twitter"
	"github.com/danbondd/tweet/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "not enough arguments passed to command")
		os.Exit(-1)
	}

	c, err := config.New(utils.FileReader, utils.NewJSONDecoder)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}

	tweet := twitter.NewTweet(http.DefaultClient, http.NewRequest)
	res, err := tweet.Send(c, os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error sending tweet: %v\n", err)
		os.Exit(-1)
	}

	fmt.Println(res)
}
