package main

import (
	"fmt"
	"os"
	"time"
)

func commandExit(cfg *config, args ...string) error {
	err := savePokedexData(&pokedexData)
	if err != nil {
		fmt.Println("Error saving Pokedex data:", err)
	}
	fmt.Println("Saving your Pokedex...")
	time.Sleep(2500 * time.Millisecond)
	fmt.Println("Closing the Pokedex... Goodbye!")
	line.Close()
	os.Exit(0)
	return nil
}
