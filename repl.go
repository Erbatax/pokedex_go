package main

import (
	"strings"

	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Show available commands",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func startRepl() {
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
		command := getCommands()[commandName]
		if command.callback != nil {
			err := command.callback()
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
