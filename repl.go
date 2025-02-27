package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mogumogu934/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		commandName := input[0]
		args := input[1:]

		cmd, exists := commands[commandName]
		if exists {
			err := cmd.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	return strings.Fields(lower)
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},

		"mapf": {
			name:        "mapf",
			description: "Displays the next page of locations",
			callback:    commandMapf,
		},

		"mapb": {
			name:        "mapb",
			description: "Displays the previous page of locations",
			callback:    commandMapb,
		},

		"explore": {
			name:        "explore <location ID | location name>",
			description: "Displays list of Pokemon in location",
			callback:    commandExplore,
		},

		"catch": {
			name:        "catch <pokemon name> <ball type>",
			description: "Attempts to catch a Pokemon",
			callback:    commandCatch,
		},

		"inspect": {
			name:        "inspect <pokemon name>",
			description: "Inspects a Pokemon",
			callback:    commandInspect,
		},

		"pokedex": {
			name:        "pokedex",
			description: "Displays list of all Pokemon caught",
			callback:    commandPokedex,
		},

		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
