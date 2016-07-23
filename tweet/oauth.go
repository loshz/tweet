package tweet

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/url"
	"strings"
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

func (oa OAuthDetails) String() string {
	return fmt.Sprintf(`OAuth
		oauth_consumer_key="%s",
		oauth_nonce="%s",
		oauth_signature="%s",
		oauth_signature_method="%s",
		oauth_timestamp="%s",
		oauth_token="%s",
		oauth_version="%s"`,
		oa.ConsumerKey, oa.Nonce, oa.Signature, signatureMethod, oa.Timestamp, oa.Token, oAuthVersion,
	)
}

// NewOAuthDetails l
func NewOAuthDetails(c *config.Config, status string) *OAuthDetails {
	oa := new(OAuthDetails)
	oa.ConsumerKey = c.ConsumerKey
	oa.Nonce = generateNonce()
	oa.SignatureMethod = signatureMethod
	oa.Timestamp = generateTimestamp()
	oa.Token = c.AccessToken
	oa.Version = oAuthVersion
	oa.generateSignature(status, c)

	return oa
	// return OAuthDetails{c.ConsumerKey, nonce, signature, signatureMethod, timestamp, c.AccessToken, oAuthVersion}
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

func (oa *OAuthDetails) generateSignature(status string, config *config.Config) {
	baseString := strings.Join([]string{"POST", url.QueryEscape(api + statusURI)}, "&")
	params := collectParams(oa, status)
	baseString = strings.Join([]string{baseString, params}, "&")
	oa.Signature = generateSignature(baseString, config.ConsumerSecret, config.AccessTokenSecret)
}

func collectParams(oa *OAuthDetails, status string) string {
	return fmt.Sprintf(`
		oauth_consumer_key=%s&
		oauth_nonce=%s&
		oauth_signature_method=%s&
		oauth_timestamp=%s&
		oauth_token=%s&
		oauth_version=%s&
		status=%s`,
		oa.ConsumerKey, oa.Nonce, signatureMethod, oa.Timestamp, oa.Token, oAuthVersion, status,
	)
}

func generateSignature(status, consumerSecret, tokenSecret string) string {
	signingKey := strings.Join([]string{consumerSecret, tokenSecret}, "&")
	mac := hmac.New(sha1.New, []byte(signingKey))
	mac.Write([]byte(status))
	signatureBytes := mac.Sum(nil)

	return base64.StdEncoding.EncodeToString(signatureBytes)
}
