package httpclient

import (
	"net/http"
	"time"
)

// Client -
type Client struct {
	httpClient http.Client
	baseUrl    string
}

// NewClient -
func NewClient(timeout time.Duration, baseUrl string) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		baseUrl: baseUrl,
	}
}
