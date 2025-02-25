package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mogumogu934/pokedex/internal/pokecache"
)

var pokemonCache *pokecache.Cache

func init() {
	pokemonCache = pokecache.NewCache(5 * time.Minute)
}

type LocationAreaResp struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetPokemon(location string) (LocationAreaResp, error) {
	url := baseURL + "/location-area" + "/" + location

	if data, exists := pokemonCache.Get(url); exists {
		var pokemon LocationAreaResp
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return LocationAreaResp{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return LocationAreaResp{}, fmt.Errorf("response failed with status code: %d and\nbody: %s", resp.StatusCode, data)
	}

	pokemon := LocationAreaResp{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return LocationAreaResp{}, err
	}

	pokemonCache.Add(url, data)

	return pokemon, nil
}
