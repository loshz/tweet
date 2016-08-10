package tweet_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/danbondd/tweet/tweet"
)

func createHttpResponse(status int) *http.Response {
	return &http.Response{
		Body:       ioutil.NopCloser(bytes.NewBufferString("body")),
		Request:    &http.Request{},
		StatusCode: status,
	}
}

func TestInvalidResponseCode(t *testing.T) {
	res := createHttpResponse(http.StatusNotFound)

	if tweet.ValidResponse(res) {
		t.Errorf("expected invalid response")
	}
}

func TestValidResponseCode(t *testing.T) {
	res := createHttpResponse(http.StatusOK)

	if !tweet.ValidResponse(res) {
		t.Errorf("expected valid response")
	}
}
