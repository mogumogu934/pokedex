package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocationAreas(pageURL *string) (LocationAreas, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreas{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreas{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, data)
	}

	locationAreas := LocationAreas{}
	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		return LocationAreas{}, err
	}

	return locationAreas, nil
}
