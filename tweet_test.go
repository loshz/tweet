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
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

type fixture struct {
	status  string
	httpReq func(method, urlStr string, body io.Reader) (*http.Request, error)
	httpDo  func(req *http.Request) (*http.Response, error)
	err     error
}

func TestTweetSend(t *testing.T) {
	testTable := make(map[string]fixture)
	testTable["TestInvalidTweetLength"] = fixture{
		status: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce consectetur dui in metus finibus, a laoreet lectus feugiat. Donec lobortis id.",
		err:    errors.New("tweet exceeds 140 character limit"),
	}
	testTable["TestNewRequestError"] = fixture{
		httpReq: func(method, urlStr string, body io.Reader) (*http.Request, error) {
			return &http.Request{}, errors.New("mock error")
		},
		err: errors.New("error building request: mock error"),
	}
	testTable["TestRequestDoError"] = fixture{
		httpReq: func(method, urlStr string, body io.Reader) (*http.Request, error) {
			return &http.Request{
				Header: make(http.Header),
			}, nil
		},
		httpDo: func(req *http.Request) (*http.Response, error) {
			return &http.Response{}, errors.New("mock error")
		},
		err: errors.New("error performing HTTP request: mock error"),
	}
	testTable["TestRequestStatusNotOK"] = fixture{
		httpReq: func(method, urlStr string, body io.Reader) (*http.Request, error) {
			return &http.Request{
				Header: make(http.Header),
			}, nil
		},
		httpDo: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				Body:       ioutil.NopCloser(bytes.NewBufferString(``)),
				Status:     "401 unauthorized",
				StatusCode: http.StatusUnauthorized,
			}, nil
		},
		err: errors.New("401 unauthorized"),
	}
	testTable["TestSuccess"] = fixture{
		httpReq: func(method, urlStr string, body io.Reader) (*http.Request, error) {
			return &http.Request{
				Header: make(http.Header),
			}, nil
		},
		httpDo: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				Body:       ioutil.NopCloser(bytes.NewBufferString(``)),
				StatusCode: http.StatusOK,
			}, nil
		},
	}

	for name, test := range testTable {
		t.Run(name, func(t *testing.T) {
			client := mockHTTPClient{
				DoFunc: test.httpDo,
			}
			tweet := NewTweet(client, test.httpReq)
			err := tweet.Send(&Config{}, test.status)
			if test.err != nil && test.err.Error() != err.Error() {
				t.Errorf("expected error: '%v', got: '%v'", test.err, err)
			}
			if test.err == nil && err != nil {
				t.Errorf("expected error: nil, got: %v", err)
			}
		})
	}
}
