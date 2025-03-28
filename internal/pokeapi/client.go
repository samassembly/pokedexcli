package pokeapi

import (
	"net/http"
	"time"
	"github.com/samassembly/pokedexcli/internal/pokecache"
)

// Client -
type Client struct {
	cache	pokecache.Cache
	httpClient http.Client
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
