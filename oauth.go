package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	apiURL          string = "https://api.twitter.com/1.1/statuses/update.json"
	signatureMethod string = "HMAC-SHA1"
	oAuthVersion    string = "1.0"
	authHeader      string = "Authorization"
)

// OAuthDetails contains a valid set of OAuth details based on credentials from a config file.
type OAuthDetails struct {
	ConsumerKey,
	Nonce,
	Signature,
	SignatureMethod,
	Timestamp,
	Token,
	Version string
}

func (oa *OAuthDetails) generateSignature(status string, c *Config) {
	params := fmt.Sprintf(`oauth_consumer_key=%s&oauth_nonce=%s&oauth_signature_method=%s&oauth_timestamp=%s&oauth_token=%s&oauth_version=%s&status=%s`,
		oa.ConsumerKey,
		oa.Nonce,
		signatureMethod,
		oa.Timestamp,
		oa.Token,
		oAuthVersion,
		encodeStatus(status),
	)

	baseString := fmt.Sprintf("%s&%s&%s", http.MethodPost, url.QueryEscape(apiURL), encodeStatus(params))
	sig := sign(baseString, c)
	oa.Signature = url.QueryEscape(sig)
}

func (oa OAuthDetails) String() string {
	return fmt.Sprintf(`OAuth oauth_consumer_key="%s", oauth_nonce="%s", oauth_signature="%s", oauth_signature_method="%s", oauth_timestamp="%s", oauth_token="%s", oauth_version="%s"`,
		oa.ConsumerKey,
		oa.Nonce,
		oa.Signature,
		signatureMethod,
		oa.Timestamp,
		oa.Token,
		oAuthVersion,
	)
}

// NewOAuthDetails collects a valid set of OAuth details based on credentials passed from a config file.
func NewOAuthDetails(c *Config, status string) *OAuthDetails {
	oa := new(OAuthDetails)
	oa.ConsumerKey = c.ConsumerKey
	oa.Nonce = generateNonce()
	oa.SignatureMethod = signatureMethod
	oa.Timestamp = fmt.Sprintf("%d", time.Now().Unix())
	oa.Token = c.AccessToken
	oa.Version = oAuthVersion
	oa.generateSignature(status, c)
	return oa
}

func generateNonce() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	b := make([]byte, 32)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func sign(baseString string, c *Config) string {
	signingKey := fmt.Sprintf("%s&%s", url.QueryEscape(c.ConsumerSecret), url.QueryEscape(c.AccessTokenSecret))
	h := hmac.New(sha1.New, []byte(signingKey))
	h.Write([]byte(baseString))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func encodeStatus(status string) string {
	return strings.Replace(url.QueryEscape(status), "+", "%20", -1)
}
