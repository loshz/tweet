package main

import (
	"fmt"
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

	res, err := tweet.Send(c, os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error sending tweet: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(res)
}
