package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New(usageError("inspect"))
	}

	pokemon := args[0]

	_, err := cfg.pokeapiClient.GetPokemonInfo(pokemon)
	if err != nil {
		return err
	}

	pokemonInfo, exists := pokedexData.PokemonMap[pokemon]
	if !exists {
		return fmt.Errorf("you have yet to catch %s", pokemon)
	}

	height := float64(pokemonInfo.Height) / 10.0
	weight := float64(pokemonInfo.Weight) / 10.0

	fmt.Printf("Name: %s\n", pokemonInfo.Name)
	fmt.Printf("Number: %d\n", pokemonInfo.ID)
	fmt.Printf("Height: %v m\n", height)
	fmt.Printf("Weight: %v kg\n", weight)
	fmt.Printf("Stats:\n")

	for _, stat := range pokemonInfo.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("Types:\n")
	for _, pType := range pokemonInfo.Types {
		fmt.Printf("  - %s\n", pType.Type.Name)
	}

	return nil
}
