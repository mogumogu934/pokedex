package main

import (
	"fmt"
)

var commandOrder = []string{
	"help",
	"mapf",
	"mapb",
	"explore",
	"exit",
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, commandName := range commandOrder {
		if cmd, exists := commands[commandName]; exists {
			fmt.Printf("%s: %s\n", commandName, cmd.description)
		}
	}
	fmt.Println()
	return nil
}
