package main

import (
	"errors"
	"fmt"
	"os"
)

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMapf(cfg *config) error {
	locations, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	fmt.Println()

	return nil
}

func commandMapb(cfg *config) error {
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
		fmt.Println(location.Name)
	}
	fmt.Println()

	return nil
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
