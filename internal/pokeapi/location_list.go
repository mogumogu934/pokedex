package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mogumogu934/pokedex/internal/pokecache"
)

type LocationAreaList struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var locationCache *pokecache.Cache

func init() {
	locationCache = pokecache.NewCache(5 * time.Minute)
}

func (c *Client) GetLocationAreas(pageURL *string) (LocationAreaList, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if data, exists := locationCache.Get(url); exists {
		var locationAreas LocationAreaList
		err := json.Unmarshal(data, &locationAreas)
		if err != nil {
			return LocationAreaList{}, err
		}
		return locationAreas, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaList{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaList{}, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaList{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return LocationAreaList{}, fmt.Errorf("response failed with status code: %d and\nbody: %s", resp.StatusCode, data)
	}

	locationAreas := LocationAreaList{}
	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		return LocationAreaList{}, err
	}

	locationCache.Add(url, data)

	return locationAreas, nil
}
