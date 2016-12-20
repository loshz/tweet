package helpers_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/danbondd/tweet/helpers"
)

func createHttpResponse(status int) *http.Response {
	return &http.Response{
		Body:       ioutil.NopCloser(bytes.NewBufferString("body")),
		Request:    &http.Request{},
		StatusCode: status,
	}
}

func TestValidResponse(t *testing.T) {
	t.Run("TestValidResponseCode", func(t *testing.T) {
		res := createHttpResponse(http.StatusNotFound)

		if helpers.ValidResponse(res) {
			t.Errorf("expected invalid response")
		}
	})

	t.Run("TestValidResponseCode", func(t *testing.T) {
		res := createHttpResponse(http.StatusOK)

		if !helpers.ValidResponse(res) {
			t.Errorf("expected valid response")
		}
	})
}
