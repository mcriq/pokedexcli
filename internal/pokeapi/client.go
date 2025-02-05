package pokeapi

import (
	"net/http"
	"time"

	"github.com/mcriq/pokedexcli/internal/pokecache"
)

// Client
type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}