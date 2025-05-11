package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL = "https://pokeapi.co/api/v2/"
)

func (c * Client) GetLocationAreas(url *string) (LocationAreas, error) {
	fullURL := baseURL + "location-area/"
	if url != nil {
		fullURL = *url
	}

	if cached, ok := c.cache.Get(fullURL); ok {
		var locationAreas LocationAreas
		err := json.Unmarshal(cached, &locationAreas)
		if err != nil {
			return LocationAreas{}, err
		}
		return locationAreas, nil
	}
	// fmt.Printf("Fetching data from API: %s\n", fullURL)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreas{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		return LocationAreas{}, fmt.Errorf("error: %s", resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreas{}, err
	}
	c.cache.Add(fullURL, data)
	var locationAreas LocationAreas
	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		return LocationAreas{}, err
	}
	return locationAreas, nil
}

func (c *Client) ExploreArea(area string) (Location, error) {
	fullURL := baseURL + "location-area/" + area
	if cached, ok := c.cache.Get(fullURL); ok {
		var location Location
		err := json.Unmarshal(cached, &location)
		if err != nil {
			return location,nil
		}
	}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Location{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		return Location{}, fmt.Errorf("error: %s", resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}
	c.cache.Add(fullURL, data)
	var location Location
	err = json.Unmarshal(data, &location)
	if err != nil {
		return Location{}, err
	}
	return location, nil
}

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	fullURL := baseURL + "pokemon/" + name
	if cached, ok := c.cache.Get(fullURL); ok {
		var pokemon Pokemon
		err := json.Unmarshal(cached, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("error: %s", resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(fullURL, data)
	var pokemon Pokemon
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}