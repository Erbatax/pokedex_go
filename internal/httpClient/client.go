package httpClient

import (
	"net/http"
	"strings"
	"time"

	"github.com/erbatax/pokedex_go/internal/httpCache"
)

// Client -
type Client struct {
	httpClient http.Client
	baseUrl    string
	cache      httpCache.Cache
}

type ClientConfig struct {
	Timeout       time.Duration
	BaseUrl       string
	CacheInterval time.Duration
}

// NewClient -
func NewClient(config ClientConfig) Client {
	return Client{
		httpClient: http.Client{
			Timeout: config.Timeout,
		},
		baseUrl: config.BaseUrl,
		cache:   *httpCache.NewCache(config.CacheInterval),
	}
}

func (c *Client) constructUrl(url *string) string {
	requestUrl := c.baseUrl
	if url != nil {
		if strings.HasPrefix(*url, c.baseUrl) == false {
			requestUrl += *url
		} else {
			requestUrl = *url
		}
	}

	return requestUrl
}
