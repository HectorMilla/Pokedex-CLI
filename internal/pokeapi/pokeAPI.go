package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/HectorMilla/Pokedex-CLI/internal/types"
	"net/http"
)

const baseURL string = "https://pokeapi.co/api/v2"

func (c *Client) GetPokeLocations(pageUrl *string) (types.LocationArea, error) {
	var locArea types.LocationArea
	url := baseURL + "/location-area"
	if value, ok := c.cache.Get(*pageUrl); ok {

		locArea = value

	} else {
		if *pageUrl != "" {
			url = *pageUrl
		}
		res, err := http.Get(url)

		if err != nil {
			fmt.Printf("Error: %v", err)
			return types.LocationArea{}, fmt.Errorf("error fetching data: %w", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return types.LocationArea{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
		}

		if err := json.NewDecoder(res.Body).Decode(&locArea); err != nil {
			return types.LocationArea{}, fmt.Errorf("error decoding json: %w", err)
		}

	}
	c.cache.Add(url, locArea)

	return locArea, nil
}
