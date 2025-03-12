package httputils

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// HTTPClient struct
type HTTPClient http.Client

// NewHTTPClient creates a new HTTP client instance
func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		Transport: &http.Transport{
			MaxIdleConns:        10,
			IdleConnTimeout:     30 * time.Second,
			DisableKeepAlives:   true,
			TLSHandshakeTimeout: 10 * time.Second,
		},
	}
}

// SendRequest sends an HTTP request
func (c *HTTPClient) SendRequest(method string, url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := http.Client(*c)
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return nil, err
	}

	return response, nil
}
