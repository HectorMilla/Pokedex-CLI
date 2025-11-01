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
			err := command.callback(cfg)
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

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
	}
}
