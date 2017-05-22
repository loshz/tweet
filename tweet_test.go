package main

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
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

func newConfig() *Config {
	return new(Config)
}

func TestInvalidTweetLength(t *testing.T) {
	status := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce consectetur dui in metus finibus, a laoreet lectus feugiat. Donec lobortis id."

	tweet := NewTweet(http.DefaultClient, http.NewRequest)
	err := tweet.Send(newConfig(), status)
	if err == nil {
		t.Errorf("expected error, got: nil")
	}
}

func TestSuccessfulTweet(t *testing.T) {
	status := "SUCCESS :)"

	m := mockHTTPClient{false, http.StatusOK}
	tweet := NewTweet(m, http.NewRequest)
	err := tweet.Send(newConfig(), status)
	if err != nil {
		t.Errorf("expected nil, got: %v", err)
	}
}

func TestNewRequestError(t *testing.T) {
	status := "Cannot create new request :("

	tweet := NewTweet(http.DefaultClient, mockNewRequest)
	err := tweet.Send(newConfig(), status)
	if err == nil {
		t.Errorf("expected error, got: nil")
	}
}

func TestRequestDoError(t *testing.T) {
	status := "Cannot create perform request :("

	m := mockHTTPClient{true, http.StatusInternalServerError}
	tweet := NewTweet(m, http.NewRequest)
	err := tweet.Send(newConfig(), status)
	if err == nil {
		t.Errorf("expected error, got: nil")
	}
}

func TestRequestStatusNotOK(t *testing.T) {
	status := "Bad HTTP response status code :("

	m := mockHTTPClient{false, http.StatusUnauthorized}
	tweet := NewTweet(m, http.NewRequest)
	err := tweet.Send(newConfig(), status)
	if err == nil {
		t.Errorf("expected error, got: nil")
	}
}
