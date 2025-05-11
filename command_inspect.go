package main

import (
	"fmt"
)
func commandInspect(config *Config) error {
	if len(config.args) == 0 {
		return fmt.Errorf("please provide a Pokémon name")
	}
	pokemonName := config.args[0]
	pokemon, ok := config.pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("pokemon %s not found in your Pokedex", pokemonName)
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v m\n", pokemon.Height)
	fmt.Printf("Weight: %v kg\n", pokemon.Weight)
	fmt.Printf("Stats: \n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types: \n")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil
}