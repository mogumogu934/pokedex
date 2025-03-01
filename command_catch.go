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

	if len(pokemonInLocation) == 0 {
		return errors.New("you must explore a location first")
	}

	pokemonTarget := args[0]

	pokemonInfo, err := cfg.pokeapiClient.GetPokemonInfo(pokemonTarget)
	if err != nil {
		return err
	}

	if _, exists := pokemonInLocation[pokemonTarget]; !exists {
		return fmt.Errorf("%s is not in current location", pokemonTarget)
	}

	ball := "poke-ball"
	ballRate := 100

	if len(args) == 2 {
		ball = args[1]
		ballRate = getBallRate(ball)
		if ballRate == 0 {
			return errors.New("you must use a valid ball type: 'poke-ball(default value)', 'great-ball', 'ultra-ball', 'master-ball'")
		}
	}

	catchRate := ballRate - (pokemonInfo.BaseExperience / 2)
	if catchRate > 100 {
		catchRate = 100 // Max = 100%
	}
	if catchRate < 5 {
		catchRate = 5 // Min = 5%
	}

	if ball == "ultra-ball" {
		fmt.Printf("Throwing an %s at %s...\n", ball, pokemonTarget)
	} else {
		fmt.Printf("Throwing a %s at %s...\n", ball, pokemonTarget)
	}

	// fmt.Printf("Catch rate: %d%%\n", catchRate)

	if caught := catchRate >= rand.Intn(100); caught {
		fmt.Printf("%s was caught!\n", pokemonTarget)
		pokedex[pokemonTarget] = pokemonInfo                 // For commandInspect
		pokemonCaught = append(pokemonCaught, pokemonTarget) // To handle multiple of same Pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemonTarget)
	}

	return nil
}

func getBallRate(ball string) int {
	switch ball {
	case "master-ball":
		return 500
	case "ultra-ball":
		return 160
	case "great-ball":
		return 130
	case "poke-ball":
		return 100
	default:
		return 0
	}
}
