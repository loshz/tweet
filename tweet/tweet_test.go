package tweet_test

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/danbondd/tweet/config"
	"github.com/danbondd/tweet/tweet"
)

type mockHTTPClient struct {
	doErr      bool
	statusCode int
}

func (m mockHTTPClient) Do(req *http.Request) (resp *http.Response, err error) {
	if m.doErr {
		return resp, errors.New("error performing request")
	}

	return &http.Response{
		Body:       ioutil.NopCloser(bytes.NewBufferString("body")),
		StatusCode: m.statusCode,
	}, nil
}

func mockNewRequest(method, urlStr string, body io.Reader) (*http.Request, error) {
	return &http.Request{}, errors.New("error creating request")
}

func newConfig() *config.Config {
	return new(config.Config)
}

func TestInvalidTweetLength(t *testing.T) {
	status := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce consectetur dui in metus finibus, a laoreet lectus feugiat. Donec lobortis id."

	tw := tweet.New(http.DefaultClient, http.NewRequest)
	_, err := tw.Send(newConfig(), status)
	if err == nil {
		t.Errorf("expected error, got: nil")
	}
}

func TestSuccessfulTweet(t *testing.T) {
	status := "SUCCESS :)"

	m := mockHTTPClient{false, http.StatusOK}
	tw := tweet.New(m, http.NewRequest)
	_, err := tw.Send(newConfig(), status)
	if err != nil {
		t.Errorf("expected nil, got: %v", err)
	}
}

func TestNewRequestError(t *testing.T) {
	status := "Cannot create new request :("

	tw := tweet.New(http.DefaultClient, mockNewRequest)
	_, err := tw.Send(newConfig(), status)
	if err == nil {
		t.Errorf("expected error, got: nil")
	}
}

func TestRequestDoError(t *testing.T) {
	status := "Cannot create perform request :("

	m := mockHTTPClient{true, http.StatusInternalServerError}
	tw := tweet.New(m, http.NewRequest)
	_, err := tw.Send(newConfig(), status)
	if err == nil {
		t.Errorf("expected error, got: nil")
	}
}

func TestRequestStatusNotOK(t *testing.T) {
	status := "Bad HTTP response status code :("

	m := mockHTTPClient{false, http.StatusUnauthorized}
	tw := tweet.New(m, http.NewRequest)
	_, err := tw.Send(newConfig(), status)
	if err == nil {
		t.Errorf("expected error, got: nil")
	}
}
