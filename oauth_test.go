package tweet_test

import (
	"fmt"
	"testing"

	"github.com/danbondd/tweet"
)

func TestStringOutput(t *testing.T) {
	expected := `OAuth oauth_consumer_key="ConsumerKey", oauth_nonce="Nonce", oauth_signature="Signature", oauth_signature_method="HMAC-SHA1", oauth_timestamp="Timestamp", oauth_token="Token", oauth_version="1.0"`

	oa := tweet.OAuthDetails{
		ConsumerKey: "ConsumerKey",
		Nonce:       "Nonce",
		Signature:   "Signature",
		Timestamp:   "Timestamp",
		Token:       "Token",
	}

	if fmt.Sprintf("%s", oa) != expected {
		t.Errorf("format does not match expected output")
	}
}
