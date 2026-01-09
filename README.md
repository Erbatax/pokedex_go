# Pokedex Go

A command-line interface (CLI) application built in Go that allows users to explore Pokemon locations, catch Pokemon, and manage a personal Pokedex using the PokeAPI.

This project was created as part of the [Boot.dev course: Build Pokedex CLI Golang](https://www.boot.dev/courses/build-pokedex-cli-golang).

## Features

- **Explore Locations**: Navigate through Pokemon locations using `map` and `mapb` commands.
- **Catch Pokemon**: Attempt to catch Pokemon by name with the `catch` command.
- **Inspect Pokemon**: View detailed information about caught Pokemon using `inspect`.
- **Pokedex Management**: List all caught Pokemon with the `pokedex` command.
- **Help System**: Get help on available commands with `help`.

## Installation

1. Ensure you have Go installed (version 1.25.5 or later).
2. Clone the repository:
   ```bash
   git clone https://github.com/erbatax/pokedex_go.git
   cd pokedex_go
   ```
3. Build the application:
   ```bash
   go build -o pokedex
   ```

## Usage

Run the application:

```bash
./pokedex
```

You will enter an interactive REPL. Type commands to interact with the Pokedex.

### Available Commands

- `help` - Show available commands or help for a specific command
- `exit` - Exit the Pokedex
- `map` - Show 20 next locations on the map
- `mapb` - Show 20 previous locations on the map
- `explore <location-name>` - Explore a specific location by name
- `catch <pokemon-name>` - Catch a Pokemon by name
- `inspect <pokemon-name>` - Inspect a caught Pokemon by name
- `pokedex` - List all caught Pokemon

### Example Usage

```
Pokedex > map
Pokedex > explore canalave-city
Pokedex > catch pikachu
Pokedex > inspect pikachu
Pokedex > pokedex
Pokedex > exit
```

## Dependencies

This project uses the PokeAPI (https://pokeapi.co/) to fetch Pokemon data. No external Go dependencies are required beyond the standard library.

## Contributing

This is a course project, but feel free to fork and experiment!

## License

This project is for educational purposes.

## Ideas for extending the project

- Update the CLI to support the "up" arrow to cycle through previous commands
- Simulate battles between pokemon
- Add more unit tests
- Refactor your code to organize it better and make it more testable
- Keep pokemon in a "party" and allow them to level up
- Allow for pokemon that are caught to evolve after a set amount of time
- Persist a user's Pokedex to disk so they can save progress between sessions
- Use the PokeAPI to make exploration more interesting. For example, rather than typing the names of areas, maybe you are given choices of areas and just type "left" or "right"
- Random encounters with wild pokemon
- Adding support for different types of balls (Pokeballs, Great Balls, Ultra Balls, etc), which have different chances of catching pokemon
