package main

import (
	"fmt"
	"sort"
)

func init() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Type 'help' to see list of commands")
	fmt.Println()
}

func commandHelp(cfg *config, args ...string) error {
	var names []string
	for name := range commands {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		cmd := commands[name]
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()
	return nil
}
