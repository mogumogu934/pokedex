package main

import (
	"errors"
	"fmt"
	"strings"
)

func extractID(url string) string {
	url = strings.TrimSuffix(url, "/")
	parts := strings.Split(url, "/")

	return parts[len(parts)-1]
}

func commandMapf(cfg *config, args ...string) error {
	locations, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Printf("%s %s\n", extractID(location.URL), location.Name)
	}

	fmt.Println()
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locations, err := cfg.pokeapiClient.GetLocationAreas(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Printf("%s %s\n", extractID(location.URL), location.Name)
	}

	fmt.Println()
	return nil
}
