package pokeapi

import (
	"net/http"
	"time"

	"github.com/Jasperino64/pokedexcli/internal/pokecache"
)

// Client -
type Client struct {
	httpClient http.Client
	cache	  *pokecache.Cache
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}
