package main

import (
	"fmt"
)

func commandHelp(config *config, args ...string) error {
	if len(args) == 0 {
		// Show general help
		// list all commands with their descriptions
		fmt.Println("Welcome to the Pokedex!")
		fmt.Println("Usage:")
		fmt.Println("help - Show this help message")
		fmt.Println("help <command> - Show help for a specific command")
		fmt.Println()
		fmt.Println("Available commands:")
		for _, command := range getCommands() {
			fmt.Printf("%s: %s\n", command.name, command.description)
		}
		return nil
	}

	// Show help for a specific command
	commandName := args[0]
	command, exists := getCommands()[commandName]
	if !exists {
		fmt.Printf("No help available for unknown command: %s\n", commandName)
		return nil
	}

	fmt.Printf("Help for command '%s':\n", command.name)
	fmt.Printf("Description: %s\n", command.description)
	fmt.Printf("Usage: %s\n", command.usage)

	return nil
}
