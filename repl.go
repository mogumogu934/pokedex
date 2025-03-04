package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/mogumogu934/pokedex/internal/pokeapi"
	"github.com/peterh/liner"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

var line *liner.State

func startRepl(cfg *config) {
	line = liner.NewLiner()
	defer func() {
		if line != nil {
			line.Close() // Ensure terminal state is reset
		}
	}()

	setupSignalHandler()

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Type 'help' to view list of commands")
	line.AppendHistory("help")

	for {
		fmt.Println()
		rawInput, _ := line.Prompt("Pokedex > ")
		if len(rawInput) == 0 {
			continue
		}

		input := cleanInput(rawInput)
		commandName := input[0]
		args := input[1:]

		cmd, exists := commands[commandName]
		if exists {
			err := cmd.callback(cfg, args...)
			line.AppendHistory(rawInput)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("unknown command")
			continue
		}
	}
}

func setupSignalHandler() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Println("\nSignal received! Cleaning up...")
		if line != nil {
			line.Close() // Restore terminal state
		}
		os.Exit(0)
	}()
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
		"catch": {
			name:        "catch <pokemon name> <ball type>",
			description: "Attempts to catch a Pokemon",
			callback:    commandCatch,
		},

		"cry": {
			name:        "cry <pokemon name>",
			description: "Plays a Pokemon's cry",
			callback:    commandCry,
		},

		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},

		"explore": {
			name:        "explore <location ID | location name>",
			description: "Displays list of Pokemon in location",
			callback:    commandExplore,
		},

		"help": {
			name:        "help",
			description: "Displays list of commands",
			callback:    commandHelp,
		},

		"inspect": {
			name:        "inspect <pokemon name>",
			description: "Displays information about a Pokemon",
			callback:    commandInspect,
		},

		"mapb": {
			name:        "mapb",
			description: "Displays the previous page of locations",
			callback:    commandMapb,
		},

		"mapf": {
			name:        "mapf",
			description: "Displays the next page of locations",
			callback:    commandMapf,
		},

		"pokedex": {
			name:        "pokedex",
			description: "Displays list of all Pokemon caught",
			callback:    commandPokedex,
		},
	}
}

func usageError(cmd string) string {
	return fmt.Sprintf("usage: %s", commands[cmd].name)
}
