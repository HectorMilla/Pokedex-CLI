package main

import "fmt"

func commandExplore(cfg *config, additionalCommands []string) error {
	if len(additionalCommands) > 1 {
		location := additionalCommands[1]
		pokemonEncounters, _ := cfg.pokeapiClient.GetPokemonAtLocation(location)
		fmt.Printf("Exploring %v...\nFound Pokemon:\n", location)
		for _, encounter := range pokemonEncounters.PokemonEncounters {
			fmt.Printf("- %v\n", encounter.Pokemon.Name)
		}
		return nil
	} else {
		//fmt.Printf("len of add commands %v", len(additionalCommands))
		return fmt.Errorf("to use the explore command use \"explore [location name]\"")
	}
}
