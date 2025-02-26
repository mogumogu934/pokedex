package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mogumogu934/pokedex/internal/pokecache"
)

type PokemonInfo struct {
	BaseExperience         int    `json:"base_experience"`
	Height                 int    `json:"height"`
	ID                     int    `json:"id"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Name                   string `json:"name"`
	Stats                  []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

var pokemonInfoCache *pokecache.Cache

func init() {
	pokemonInfoCache = pokecache.NewCache(5 * time.Minute)
}

func (c *Client) GetPokemonInfo(pokemon string) (PokemonInfo, error) {
	url := baseURL + "/pokemon" + "/" + pokemon

	if data, exists := pokemonInfoCache.Get(url); exists {
		var pokemonInfo PokemonInfo
		err := json.Unmarshal(data, &pokemonInfo)
		if err != nil {
			return PokemonInfo{}, err
		}
		return pokemonInfo, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInfo{}, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return PokemonInfo{}, fmt.Errorf("you must provide the name of a valid pokemon")
	}

	if resp.StatusCode > 299 {
		return PokemonInfo{}, fmt.Errorf("response failed with status code: %d and\nbody: %s", resp.StatusCode, data)
	}

	pokemonInfo := PokemonInfo{}
	err = json.Unmarshal(data, &pokemonInfo)
	if err != nil {
		return PokemonInfo{}, err
	}

	pokemonInfoCache.Add(url, data)

	return pokemonInfo, nil
}
