package main

import (
	"fmt"
)

func commandPokedex(config *Config) error {
	fmt.Println("Your Pokédex:")
	if len(config.pokedex) == 0 {
		fmt.Println("  No Pokémon caught yet.")
	} else {
		for name, _ := range config.pokedex {
			fmt.Printf("  - %s\n", name)
		}
	}
	return nil
}