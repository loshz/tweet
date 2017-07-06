package main

import (
	"fmt"
	"testing"
)

func TestStringOutput(t *testing.T) {
	expected := `OAuth oauth_consumer_key="ConsumerKey", oauth_nonce="Nonce", oauth_signature="Signature", oauth_signature_method="HMAC-SHA1", oauth_timestamp="Timestamp", oauth_token="Token", oauth_version="1.0"`

	oa := OAuthDetails{
		ConsumerKey: "ConsumerKey",
		Nonce:       "Nonce",
		Signature:   "Signature",
		Timestamp:   "Timestamp",
		Token:       "Token",
	}

	if fmt.Sprint(oa) != expected {
		t.Errorf("format does not match expected output")
	}
}
