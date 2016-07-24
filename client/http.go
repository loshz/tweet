package client

import "net/http"

// HTTPClient contains a HTTP that issues a POST to the specified URL.
type HTTPClient interface {
	//NewRequest(method, urlStr string, body io.Reader) (*http.Request, error)
	Do(req *http.Request) (resp *http.Response, err error)
}

// ValidResponse checks a given HTTP response is valid and doesn't contain errors.
func ValidResponse(res *http.Response) bool {
	return res.StatusCode == http.StatusOK
}
