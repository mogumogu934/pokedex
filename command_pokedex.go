package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(pokemonCaught) == 0 {
		return errors.New("you have yet to catch any Pokemon")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range pokemonCaught {
		fmt.Printf("  - %s\n", pokemon)
	}

	return nil
}
