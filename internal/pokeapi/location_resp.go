package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mogumogu934/pokedex/internal/pokecache"
)

type LocationAreaResp struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

var locationAreaRespCache *pokecache.Cache

func init() {
	locationAreaRespCache = pokecache.NewCache(5 * time.Minute)
}

func (c *Client) GetLocationAreaResp(location string) (LocationAreaResp, error) {
	url := baseURL + "/location-area" + "/" + location

	if data, exists := locationAreaRespCache.Get(url); exists {
		var locationAreaResp LocationAreaResp
		err := json.Unmarshal(data, &locationAreaResp)
		if err != nil {
			return LocationAreaResp{}, err
		}
		return locationAreaResp, nil
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

	if resp.StatusCode == 404 {
		return LocationAreaResp{}, errors.New("you must provide a valid location ID or location name")
	}

	if resp.StatusCode > 299 {
		return LocationAreaResp{}, fmt.Errorf("response failed with status code: %d and\nbody: %s", resp.StatusCode, data)
	}

	locationAreaResp := LocationAreaResp{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationAreaResp{}, err
	}

	locationAreaRespCache.Add(url, data)

	return locationAreaResp, nil
}
