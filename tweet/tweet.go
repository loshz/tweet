package tweet

import "github.com/danbondd/tweet/config"

const (
	tweetLength int = 140
)

// Send l
func Send(c *config.Config, tweet string) (string, error) {
	_ = NewOAuthDetails(c, tweet)

	return tweet, nil
}
