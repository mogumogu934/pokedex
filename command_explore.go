package main

import (
	"errors"
	"fmt"
)

var PokemonInLocation map[string]int

func init() {
	PokemonInLocation = make(map[string]int)
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you must provide a location ID or location name")
	}
	clear(PokemonInLocation) // Delete all elements from map
	locationArea := args[0]

	locationAreaResp, err := cfg.pokeapiClient.GetLocationAreaResp(locationArea)
	if err != nil {
		return err
	}

	pokemon := locationAreaResp.PokemonEncounters

	fmt.Printf("Exploring %s...\n", locationAreaResp.Name)
	fmt.Println("Found Pokemon:")
	for _, p := range pokemon {
		fmt.Printf("- %s\n", p.Pokemon.Name)
		PokemonInLocation[p.Pokemon.Name] = 0
	}

	fmt.Println()
	return nil
}
