package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/i-bielik/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	Next          *string `json:"next"`
	Previous      *string `json:"previous"`
}

func cleanInput(text string) []string {
	// Convert to lowercase and split into words
	return strings.Fields(strings.ToLower(text))
}

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex")
	return nil
}

func commandMap(cfg *config, args ...string) error {
	locations, err := cfg.pokeapiClient.GetLocationAreas(cfg.Next)
	if err != nil {
		return err
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return errors.New("you're on the first page")
	}

	locations, err := cfg.pokeapiClient.GetLocationAreas(cfg.Previous)
	if err != nil {
		return err
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("please provide a location name to explore")
	}
	locationName := args[0]
	location, err := cfg.pokeapiClient.GetLocationArea(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationName)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range location.PokemonEncounters {
		fmt.Println("-", pokemon.Pokemon.Name)
	}
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func startRepl(cfg *config) {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Help with the Pokedex",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get previous location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location name>",
			description: "Explore Pokemons in given location",
			callback:    commandExplore,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]

		if command, ok := commands[commandName]; ok {
			command.callback(cfg, words[1:]...)
		} else {
			fmt.Println("Unknown command:", commandName)
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			return
		}
	}

}
