package tweet

import (
	"bytes"
	"math/rand"
	"time"

	"github.com/danbondd/tweet/config"
)

const (
	api             string = "https://api.twitter.com/"
	apiVersion      string = "1.1"
	statusURI       string = "/statuses/update.json"
	signatureMethod string = "HMAC-SHA1"
	oAuthVersion    string = "1.0"
)

// OAuthDetails l
type OAuthDetails struct {
	ConsumerKey     string `json:"oauth_consumer_key"`
	Nonce           string `json:"oauth_nonce"`
	Signature       string `json:"oauth_signature"`
	SignatureMethod string `json:"oauth_signature_method"`
	Timestamp       string `json:"oauth_timestamp"`
	Token           string `json:"oauth_token"`
	Version         string `json:"oauth_version"`
}

// NewOAuthDetails l
func NewOAuthDetails(c *config.Config, status string) OAuthDetails {
	nonce := generateNonce()
	timestamp := generateTimestamp()
	signature := generateSignature(status)

	return OAuthDetails{c.ConsumerKey, nonce, signature, signatureMethod, timestamp, c.AccessToken, oAuthVersion}
}

func generateNonce() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	b := make([]byte, 32)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}

	return string(b)
}

func generateTimestamp() string {
	s := time.Now().UnixNano() / int64(time.Millisecond)
	return string(s)
}

func generateSignature(params string) string {
	var buff bytes.Buffer
	buff.WriteString("status=" + params)
	return ""
}
