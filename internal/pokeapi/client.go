package pokeapi

import (
	"time"

	"github.com/erbatax/pokedex_go/internal/httpClient"
)

// Client -
type Client struct {
	pokeApi httpClient.Client
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	return Client{
		pokeApi: httpClient.NewClient(httpClient.ClientConfig{
			Timeout:       timeout,
			BaseUrl:       "https://pokeapi.co/api/v2",
			CacheInterval: 10 * time.Minute,
		}),
	}
}
