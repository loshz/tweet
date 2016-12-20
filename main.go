package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/danbondd/tweet/config"
	"github.com/danbondd/tweet/helpers"
	"github.com/danbondd/tweet/twitter"
)

func main() {
	c, err := config.New(helpers.FileReader, helpers.NewJSONDecoder)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	tweet := twitter.NewTweet(http.DefaultClient, http.NewRequest)
	res, err := tweet.Send(c, os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error sending tweet: %v\n", err)
		return
	}

	fmt.Println(res)
}
