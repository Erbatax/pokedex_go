package pokeapi

import (
	httpclient "github.com/erbatax/pokedex_go/internal/httpClient"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	return httpclient.Get[RespShallowLocations](&c.pokeApi, &url)
}
