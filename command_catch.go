package main

import (
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, additionalCommands []string) error {
	if len(additionalCommands) <= 1 {
		return fmt.Errorf("invalid command to catch a pokemon us \"catch <pokemon name>\"")
	}
	pokemonName := additionalCommands[1]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)
	time.Sleep(500 * time.Millisecond)
	randomNumber := rand.Intn(800) + 1
	if pokemon.BaseExperience <= randomNumber {
		fmt.Printf("%v was caught!\n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.pokeapiClient.AddPokemonToPokedex(pokemon)
	} else {
		fmt.Printf("%v escaped!\n", pokemon.Name)
	}

	return nil
}
