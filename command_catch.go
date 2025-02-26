package main

import (
	"fmt"
	"math/rand"

	"github.com/mogumogu934/pokedex/internal/pokeapi"
)

var pokedex = make(map[string]pokeapi.PokemonInfo)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide the name of a pokemon")
	}

	pokemon := args[0]

	pokemonInfo, err := cfg.pokeapiClient.GetPokemonInfo(pokemon)
	if err != nil {
		return err
	}

	catchRate := 100 - (pokemonInfo.BaseExperience / 3)
	if catchRate < 10 {
		catchRate = 10 // Minimum is 10%
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	if caught := catchRate >= rand.Intn(100); caught {
		fmt.Printf("%s was caught!\n", pokemon)
		pokedex[pokemon] = pokemonInfo
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}

	fmt.Println()
	return nil
}
