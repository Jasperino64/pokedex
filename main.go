package main

import (
	"time"

	"github.com/Jasperino64/pokedexcli/internal/pokeapi"
)
func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	config := &Config{
		client:   &pokeClient,
		pokedex: make(map[string]pokeapi.Pokemon),
	}
	startRepl(config)
}