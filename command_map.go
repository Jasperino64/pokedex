package main

import (
	"fmt"
)

func commandMap(config *Config) error {
	locationAreas, err := config.client.GetLocationAreas(config.nextUrl)
	if err != nil {
		return err
	}
	config.prevUrl = config.nextUrl
	config.nextUrl = locationAreas.Next
	if locationAreas.Previous != nil {
		config.prevUrl = locationAreas.Previous
	}

	for _, area := range locationAreas.Results {
		if area.Name == "" {
			continue
		}
		fmt.Printf("%s\n", area.Name)
	}
	return nil
}

func commandMapb(config *Config) error {
	if config.prevUrl == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	locationAreas, err := config.client.GetLocationAreas(config.prevUrl)
	if err != nil {
		return err
	}
	if locationAreas.Previous == nil {
		config.prevUrl = nil
	} else {
		config.prevUrl = locationAreas.Previous
	}
	config.nextUrl = locationAreas.Next
	if locationAreas.Previous != nil {
		config.prevUrl = locationAreas.Previous
	}

	for _, area := range locationAreas.Results {
		if area.Name == "" {
			continue
		}
		fmt.Printf("%s\n", area.Name)
	}
	return nil
}