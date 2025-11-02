package main

import (
	"bufio"
	"fmt"
	"os"
	s "strings"

	"github.com/HectorMilla/Pokedex-CLI/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	next          string
	previous      string
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, additionalCommands []string) error
}

func StartRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex >")
		scanner.Scan()

		input := CleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		commandName := input[0]
		command, exist := getCommands()[commandName]

		if exist {
			err := command.callback(cfg, input)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

func CleanInput(text string) []string {
	inputText := s.ToLower(text)
	listOfWords := s.Fields(inputText)
	return listOfWords
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Show pokedex usage",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "List of pokemon location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous pokemon locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Shows list of pokemon located at given location using \"explore <location name>\"",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "catch a pokemon using \"catch <pokemon name>\"",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "get more information on a pokemon you have caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "displays a list of pokemon added to your pokedex",
			callback:    commandPokedex,
		},
	}
}
