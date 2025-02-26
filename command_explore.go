package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide a location ID or location name")
	}

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
	}

	fmt.Println()
	return nil
}
