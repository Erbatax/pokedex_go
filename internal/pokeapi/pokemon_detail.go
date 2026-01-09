package pokeapi

import (
	"fmt"
	"strings"

	"github.com/erbatax/pokedex_go/internal/httpClient"
)

// ListLocations -
func (c *Client) GetPokemonDetail(pokemonName string) (ResponsePokemon, error) {
	url := "/pokemon/" + pokemonName

	pokemon, err := httpClient.Get[ResponsePokemon](&c.pokeApi, &url)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			var zero ResponsePokemon
			return zero, fmt.Errorf("Pokemon %s not found.\n", pokemonName)
		}
		var zero ResponsePokemon
		return zero, err
	}

	return pokemon, nil
}
