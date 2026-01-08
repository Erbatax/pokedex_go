package pokeapi

import (
	"time"

	httpclient "github.com/erbatax/pokedex_go/internal/httpClient"
)

// Client -
type Client struct {
	pokeApi httpclient.Client
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	return Client{
		pokeApi: httpclient.NewClient(timeout, baseURL),
	}
}
