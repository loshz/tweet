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

	}

	res, err := tweet.Send(c, os.Args[1])
	if err != nil {

	}

	fmt.Println(res)
}
