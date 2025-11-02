package main

import "fmt"

func commandInspect(cfg *config, additionalCommands []string) error {
	if len(additionalCommands) <= 1 {
		return fmt.Errorf("to inspect a pokemon use \"inspect <pokemon name>\"")
	}
	pokemonName := additionalCommands[1]
	pokemon, err := cfg.pokeapiClient.UserPokedex.Pokemon[pokemonName]
	if !err {
		return fmt.Errorf("this pokemon does not exist in your pokedex. To add a pokemon to the pokedex you must cath it first")
	}
	fmt.Printf("Name: %v\nHeight: %v\nWeight: %v\nStats:\n", pokemon.Name, pokemon.Height, pokemon.Weight)
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("  - %v\n", pokemonType.Type.Name)
	}
	return nil
}
