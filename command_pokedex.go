package main

import "fmt"

func commandPokedex(cfg *config, additionalCommands []string) error {
	listOfPokemon := cfg.pokeapiClient.UserPokedex.Pokemon
	if len(listOfPokemon) == 0 {
		return fmt.Errorf("No pokemon caught. Catch pokemon to add them to your pokedex")
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range listOfPokemon {
		fmt.Printf("  - %v\n", pokemon.Name)
	}
	return nil
}
