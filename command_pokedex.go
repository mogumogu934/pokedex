package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	if len(pokemonCaught) == 0 {
		fmt.Println("you have yet to catch any Pokemon...")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range pokemonCaught {
		fmt.Printf("  - %s\n", pokemon)
	}

	fmt.Println()
	return nil
}
