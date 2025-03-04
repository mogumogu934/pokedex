package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(pokedexData.PokemonList) == 0 {
		return errors.New("you have yet to catch any Pokemon")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range pokedexData.PokemonList {
		fmt.Printf("  - %s\n", pokemon)
	}

	return nil
}
