package client_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/danbondd/tweet/client"
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

	if client.ValidResponse(res) {
		t.Errorf("expected error")
	}
}
