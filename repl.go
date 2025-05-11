package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Jasperino64/pokedexcli/internal/pokeapi"
)
type cliCommand struct {
	name        string
	description string
	callback    func(config * Config) error
}

type Config struct {
	client *pokeapi.Client
	prevUrl *string
	nextUrl *string
	args []string
	pokedex map[string]pokeapi.Pokemon
}

func startRepl(config *Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			if len(words) > 1 {
				config.args = words[1:]
			}
			err := command.callback(config)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the namees of 20 locations. Each subsequent call will display the next 20 locations.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the namees of 20 locations. Each subsequent call will display the previous 20 locations.",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explores a location area and displays the Pokemon that can be found there.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catches a Pokemon.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspects a Pokemon.",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays your Pokedex.",
			callback:    commandPokedex,
		},
	}
}
