package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/HectorMilla/Pokedex-CLI/internal/types"
)

const baseURL string = "https://pokeapi.co/api/v2"

func (c *Client) GetPokeLocations(pageUrl *string) (types.LocationArea, error) {
	var locArea types.LocationArea
	url := baseURL + "/location-area"
	if value, ok := c.cache.Get(*pageUrl); ok {
		err := json.Unmarshal(value, &locArea)
		if err != nil {
			return locArea, fmt.Errorf("error decoding json: %w", err)
		}
		return locArea, nil

	} else {
		if *pageUrl != "" {
			url = *pageUrl
		}
		res, err := http.Get(url)

		if err != nil {
			return types.LocationArea{}, fmt.Errorf("error fetching data: %w", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return types.LocationArea{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
		}

		results, _ := io.ReadAll(res.Body)

		if err := json.Unmarshal(results, &locArea); err != nil {
			return types.LocationArea{}, fmt.Errorf("error decoding json: %w", err)
		}
		c.cache.Add(url, results)

		return locArea, nil
	}

}

func (c *Client) GetPokemonAtLocation(location string) (types.PokemonEncounters, error) {
	var pokemonList types.PokemonEncounters
	url := baseURL + "/location-area/" + location
	if data, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(data, &pokemonList); err != nil {
			return pokemonList, fmt.Errorf("error decoding json: %v", err)
		}
		return pokemonList, nil
	}

	res, err := http.Get(url)

	if err != nil {
		return pokemonList, fmt.Errorf("error getting data from PokeApi please try again: %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return pokemonList, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	body, err := io.ReadAll(res.Body)

	if err := json.Unmarshal(body, &pokemonList); err != nil {
		return pokemonList, fmt.Errorf("error decoding json: %v", err)
	}
	c.cache.Add(url, body)
	return pokemonList, nil
}

func (c *Client) GetPokemon(pokemonName string) (types.Pokemon, error) {
	var pokemon types.Pokemon
	url := baseURL + "/pokemon/" + pokemonName

	if data, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(data, &pokemon); err != nil {
			return pokemon, fmt.Errorf("error decoding json")
		}
		return pokemon, nil
	}

	res, err := http.Get(url)

	if err != nil {
		return pokemon, fmt.Errorf("error getting pokemon information %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return pokemon, fmt.Errorf("pokemon \"%v\" not found", pokemonName)
	}
	if res.StatusCode != http.StatusOK {
		return pokemon, fmt.Errorf("unexpected http status code: %v", res.StatusCode)
	}

	body, _ := io.ReadAll(res.Body)

	if err := json.Unmarshal(body, &pokemon); err != nil {
		return pokemon, fmt.Errorf("error decoding json")
	}

	c.cache.Add(url, body)
	return pokemon, nil
}
