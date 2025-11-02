package pokeapi

import (
	"net/http"
	"time"

	"github.com/HectorMilla/Pokedex-CLI/internal/pokecache"
	"github.com/HectorMilla/Pokedex-CLI/internal/types"
)

type Client struct {
	cache       pokecache.Cache
	UserPokedex Pokedex
	httpClient  http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		UserPokedex: Pokedex{
			Pokemon: make(map[string]types.Pokemon),
		},
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
