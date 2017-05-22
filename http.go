package tweet

import (
	"io"
	"net/http"
)

// HTTPClient specifies the behaviour of a given HTTP client.
// It only implements the Do method which sends an HTTP request and returns an HTTP response.
type HTTPClient interface {
	Do(req *http.Request) (resp *http.Response, err error)
}

// NewRequest returns a new HTTP Request given a method, URL, and optional body.
type NewRequest func(method, urlStr string, body io.Reader) (*http.Request, error)
