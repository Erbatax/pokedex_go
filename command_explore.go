package main

import (
	"fmt"
)

func commandExplore(config *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Please provide a location name to explore.")
		return nil
	}

	locationName := args[0]
	if locationName == "" {
		fmt.Println("Please provide a location name to explore.")
		return nil
	}

	areaDetail, err := config.pokeapiClient.GetLocationAreaDetail(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", areaDetail.Name)
	fmt.Println("Found Pokemon:")

	for _, encounter := range areaDetail.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
