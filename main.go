package main

import (
	"time"

	"github.com/erbatax/pokedex_go/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	config := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(config)
}
