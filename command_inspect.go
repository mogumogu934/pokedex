package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide the name of a pokemon")
	}

	pokemon := args[0]
	if _, exists := pokedex[pokemon]; !exists {
		return fmt.Errorf("you have yet to catch %s", pokemon)
	}

	height := float64(pokedex[pokemon].Height) / 10.0
	weight := float64(pokedex[pokemon].Weight) / 10.0

	fmt.Printf("Name: %s\n", pokedex[pokemon].Name)
	fmt.Printf("Number: %d\n", pokedex[pokemon].ID)
	fmt.Printf("Height: %v m\n", height)
	fmt.Printf("Weight: %v kg\n", weight)
	fmt.Printf("Stats:\n")

	for _, stat := range pokedex[pokemon].Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("Types:\n")
	for _, pType := range pokedex[pokemon].Types {
		fmt.Printf("  - %s\n", pType.Type.Name)
	}

	fmt.Println()
	return nil
}
