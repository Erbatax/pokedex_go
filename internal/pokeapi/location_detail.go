package pokeapi

import (
	"github.com/erbatax/pokedex_go/internal/httpClient"
)

// ListLocations -
func (c *Client) GetLocationAreaDetail(locationName string) (ResponseLocationArea, error) {
	url := "/location-area/" + locationName

	return httpClient.Get[ResponseLocationArea](&c.pokeApi, &url)
}
