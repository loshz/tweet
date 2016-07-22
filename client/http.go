package client

import (
	"io"
	"net/http"
)

// HTTPClient contains a HTTP that issues a POST to the specified URL.
type HTTPClient interface {
	Post(url string, bodyType string, body io.Reader) (resp *http.Response, err error)
}

// ValidResponse checks a given HTTP response is valid and doesn't contain errors.
func ValidResponse(res *http.Response) bool {
	return res.StatusCode == http.StatusOK
}
