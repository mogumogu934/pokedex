package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you must provide the name of a valid pokemon")
	}

	pokemon := args[0]

	_, err := cfg.pokeapiClient.GetPokemonInfo(pokemon)
	if err != nil {
		return err
	}

	if _, exists := pokedexData.PokemonMap[pokemon]; !exists {
		return fmt.Errorf("you have yet to catch %s", pokemon)
	}

	height := float64(pokedexData.PokemonMap[pokemon].Height) / 10.0
	weight := float64(pokedexData.PokemonMap[pokemon].Weight) / 10.0

	fmt.Printf("Name: %s\n", pokedexData.PokemonMap[pokemon].Name)
	fmt.Printf("Number: %d\n", pokedexData.PokemonMap[pokemon].ID)
	fmt.Printf("Height: %v m\n", height)
	fmt.Printf("Weight: %v kg\n", weight)
	fmt.Printf("Stats:\n")

	for _, stat := range pokedexData.PokemonMap[pokemon].Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("Types:\n")
	for _, pType := range pokedexData.PokemonMap[pokemon].Types {
		fmt.Printf("  - %s\n", pType.Type.Name)
	}

	return nil
}
