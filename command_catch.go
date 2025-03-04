package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/mogumogu934/pokedex/internal/pokeapi"
)

var pokedexData = PokedexData{}

func init() {
	loadedData, err := loadPokedexData()
	if err != nil {
		pokedexData = PokedexData{
			PokemonMap:  make(map[string]pokeapi.PokemonInfo),
			PokemonList: []string{},
		}
	} else {
		pokedexData = *loadedData
	}
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New(usageError("catch"))
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

	time.Sleep(1000 * time.Millisecond)
	fmt.Println("...")
	time.Sleep(1250 * time.Millisecond)
	fmt.Println("...")
	time.Sleep(1250 * time.Millisecond)
	fmt.Println("...")
	time.Sleep(2500 * time.Millisecond)

	if caught := catchRate >= rand.Intn(100); caught {
		fmt.Printf("%s was caught!\n", pokemonTarget)
		pokedexData.PokemonMap[pokemonTarget] = pokemonInfo                      // For commandInspect
		pokedexData.PokemonList = append(pokedexData.PokemonList, pokemonTarget) // To handle multiple of same Pokemon
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
		return 170
	case "great-ball":
		return 135
	case "poke-ball":
		return 100
	default:
		return 0
	}
}

type PokedexData struct {
	PokemonMap  map[string]pokeapi.PokemonInfo `json:"pokemon_map"`
	PokemonList []string                       `json:"pokemon_list"`
}

func savePokedexData(data *PokedexData) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return os.WriteFile("pokedex.json", jsonData, 0644)
}

func loadPokedexData() (*PokedexData, error) {
	data, err := os.ReadFile("pokedex.json")
	if err != nil {
		if os.IsNotExist(err) {
			return &PokedexData{
				PokemonMap:  make(map[string]pokeapi.PokemonInfo),
				PokemonList: []string{},
			}, nil
		}
		return nil, err
	}

	var tempPokedexData PokedexData
	err = json.Unmarshal(data, &tempPokedexData)
	if err != nil {
		return nil, err
	}
	return &tempPokedexData, nil
}
