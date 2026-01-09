package main

import (
	"strings"

	"bufio"
	"fmt"
	"os"

	"github.com/erbatax/pokedex_go/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	prevLocationsUrl *string
	nextLocationsUrl *string
	pokedex          map[string]pokeapi.ResponsePokemon
}

type cliCommand struct {
	name        string
	description string
	usage       string
	callback    func(config *config, args ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Show available commands or help for a specific command",
			usage:       "help [command]",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			usage:       "exit",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Show 20 next locations on the map",
			usage:       "map",
			callback:    commandMapForward,
		},
		"mapb": {
			name:        "mapb",
			description: "Show 20 previous locations on the map",
			usage:       "mapb",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Explore a specific location by name",
			usage:       "explore <location-name>",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokémon by name",
			usage:       "catch <pokemon-name>",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a Pokémon by name",
			usage:       "inspect <pokemon-name>",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught Pokémon",
			usage:       "pokedex",
			callback:    commandPokedex,
		},
	}
}

func startRepl(config *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)
		if len(words) == 0 {
			fmt.Println("Please enter a command")
			continue
		}

		commandName := words[0]
		args := words[1:]
		command := getCommands()[commandName]
		if command.callback != nil {
			err := command.callback(config, args...)
			if err != nil {
				fmt.Printf("Error executing command: %v\n", err)
			}
		} else {
			fmt.Printf("Unknown command: %s\n", words[0])
		}
	}
}

// / Split the text into words based on whitespace, lowercase, and trim leading or trailing whitespace.
func cleanInput(text string) []string {
	words := strings.Split(text, " ")
	var cleanedWords []string
	for _, word := range words {
		trimmed := strings.TrimSpace(word)
		if trimmed != "" {
			cleanedWords = append(cleanedWords, strings.ToLower(trimmed))
		}
	}
	return cleanedWords
}
