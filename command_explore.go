package main

import (
	"fmt"
)
func commandExplore(config *Config) error {
	if len(config.args) == 0 {
		return fmt.Errorf("please provide a location area name")
	}
	fmt.Printf("Exploring %s...\n", config.args[0])
	location, err := config.client.ExploreArea(config.args[0])
	if err != nil {
		return err
	}
	if len(location.PokemonEncounters) == 0 {
		fmt.Println("No Pokémon found in this area.")
		return nil
	}
	fmt.Println("Found Pokémon:")
	for _, encounter := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}