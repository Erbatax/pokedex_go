package pokeapi

import (
	"github.com/erbatax/pokedex_go/internal/httpClient"
)

// GetLocationAreaList -
func (c *Client) GetLocationAreaList(pageURL *string) (ResponseShallowLocations, error) {
	url := "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	return httpClient.Get[ResponseShallowLocations](&c.pokeApi, &url)
}
