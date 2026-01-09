package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Please provide a pokemon name to catch.")
		return nil
	}

	pokemonName := args[0]
	if pokemonName == "" {
		fmt.Println("Please provide a pokemon name to catch.")
		return nil
	}

	pokemonDetail, err := config.pokeapiClient.GetPokemonDetail(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonDetail.Name)

	// Simulate catching the pokemon
	// A simple random chance based on base experience
	// The higher the base experience, the harder to catch
	// CatchDifficulty ranges from 0 to BaseExperience
	catchDifficulty := float64(rand.Intn(pokemonDetail.BaseExperience))
	// Pokeball effectiveness is set to 50 for basic pokeballs
	// In a real scenario, different pokeballs would have different effectiveness
	// Example: A pokemon with base experience 50 or less would have a 100% chance, but one with base experience 100 would have a 50% chance to be caught
	//
	// Base experience | Success rate
	// 50 | 100% | ■■■■■■■■■■■■■■■■■■■■
	// 100 | 50% | ■■■■■■■■■■□□□□□□□□□□
	// 150 | 33% | ■■■■■□□□□□□□□□□□□□□□
	// 200 | 25% | ■■■■□□□□□□□□□□□□□□□□
	// 250 | 20% | ■■■□□□□□□□□□□□□□□□□□
	// 300 | 16% | ■■□□□□□□□□□□□□□□□□□□
	// 350 | 14% | ■■□□□□□□□□□□□□□□□□□□
	// 400 | 12% | ■■□□□□□□□□□□□□□□□□□□
	// 450 | 11% | ■□□□□□□□□□□□□□□□□□□□
	// 500 | 10% | ■□□□□□□□□□□□□□□□□□□□

	pokeballEffectiveness := 50.0
	if catchDifficulty > pokeballEffectiveness {
		fmt.Printf("%s escaped!\n", pokemonDetail.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonDetail.Name)
	if _, exists := config.pokedex[pokemonDetail.Name]; !exists {
		fmt.Printf("%s added to your Pokedex!\n", pokemonDetail.Name)
		config.pokedex[pokemonDetail.Name] = pokemonDetail
		fmt.Println("You may now inspect it using the 'inspect' command.")
	}

	return nil
}
