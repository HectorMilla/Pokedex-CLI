package pokeapi

import "github.com/HectorMilla/Pokedex-CLI/internal/types"

type Pokedex struct {
	Pokemon map[string]types.Pokemon
}

func NewPokedex() Pokedex {
	pokedex := Pokedex{
		Pokemon: make(map[string]types.Pokemon),
	}
	return pokedex
}

func (c *Client) AddPokemonToPokedex(pokemon types.Pokemon) {
	c.UserPokedex.Pokemon[pokemon.Name] = pokemon
}
