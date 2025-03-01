package main

import (
	"fmt"
	"sort"
)

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

	return nil
}
