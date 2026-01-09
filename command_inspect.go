package main

import (
	"fmt"
)

func commandInspect(config *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Please provide a pokemon name to inspect.")
		return nil
	}

	pokemonName := args[0]
	if pokemonName == "" {
		fmt.Println("Please provide a pokemon name to inspect.")
		return nil
	}

	pokemonDetail, ok := config.pokedex[pokemonName]
	if !ok {
		fmt.Printf("You don't have %s in your Pokedex. Catch it first!\n", pokemonName)
		return nil
	}

	fmt.Printf("Name: %s\n", pokemonDetail.Name)
	fmt.Printf("Height: %.1f m\n", float64(pokemonDetail.Height)/10)
	fmt.Printf("Weight: %.1f kg\n", float64(pokemonDetail.Weight)/10)
	fmt.Printf("Base Experience: %d\n", pokemonDetail.BaseExperience)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemonDetail.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, t := range pokemonDetail.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}

	return nil
}
