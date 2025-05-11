package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *Config) error {
	if len(config.args) == 0 {
		return fmt.Errorf("please provide a Pokémon name")
	}
	pokemonName := config.args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	pokemon, err := config.client.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	catchRate := 120 - pokemon.BaseExperience
	if catchRate < 0 {
		catchRate = 10 // Minimum catch rate
	}
	roll := rand.Intn(100) + 1
	if roll <= catchRate {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		config.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}