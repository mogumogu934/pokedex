package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/mogumogu934/pokedex/internal/pokeapi"
)

var pokedex = make(map[string]pokeapi.PokemonInfo)
var pokemonCaught []string

func init() {
	pokemonCaught = make([]string, 0)
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you must provide the name of a pokemon")
	}

	if len(PokemonInLocation) == 0 {
		return errors.New("you must explore a location first")
	}

	pokemonTarget := args[0]

	if _, exists := PokemonInLocation[pokemonTarget]; !exists {
		return fmt.Errorf("%s is not in current location", pokemonTarget)
	}

	pokemonInfo, err := cfg.pokeapiClient.GetPokemonInfo(pokemonTarget)
	if err != nil {
		return err
	}

	catchRate := 100 - (pokemonInfo.BaseExperience / 2)
	if catchRate > 85 {
		catchRate = 85 // Maximum is 85%
	}
	if catchRate < 10 {
		catchRate = 10 // Minimum is 10%
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonTarget)
	if caught := catchRate >= rand.Intn(100); caught {
		fmt.Printf("%s was caught!\n", pokemonTarget)
		pokedex[pokemonTarget] = pokemonInfo                 // For commandInspect
		pokemonCaught = append(pokemonCaught, pokemonTarget) // To handle multiple of same Pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemonTarget)
	}

	fmt.Println()
	return nil
}
