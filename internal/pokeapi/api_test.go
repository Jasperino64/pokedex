package pokeapi

import (
	"testing"
	"time"
)

func TestGetLocationAreas(t *testing.T) {
	url := "https://pokeapi.co/api/v2/location-area/"
	client := NewClient(5 * time.Second, 5 * time.Minute)
	locationAreas, err := client.GetLocationAreas(&url)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if locationAreas.Count == 0 {
		t.Fatalf("Expected non-zero count, got %d", locationAreas.Count)
	}
	if len(locationAreas.Results) == 0 {
		t.Fatalf("Expected non-zero results, got %d", len(locationAreas.Results))
	}
}