package client

import (
	"io"
	"net/http"
)

// HTTPClient contains a HTTP that issues a POST to the specified URL.
type HTTPClient interface {
	Do(req *http.Request) (resp *http.Response, err error)
}

// NewRequest returns a new HTTP Request given a method, URL, and optional body.
type NewRequest func(method, urlStr string, body io.Reader) (*http.Request, error)

// ValidResponse checks a given HTTP response is valid and doesn't contain errors.
func ValidResponse(res *http.Response) bool {
	return res.StatusCode == http.StatusOK
}
