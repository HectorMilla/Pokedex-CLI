package main

import (
	"bufio"
	"fmt"
	"os"
	s "strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex >")

		for scanner.Scan() {
			result := CleanInput(scanner.Text())[0]
			if _, ok := commands[result]; ok {
				commands[result].callback()
			} else {
				fmt.Println("Unknown command")
			}
			break
		}
	}
}

func CleanInput(text string) []string {
	if len(text) == 0 {
		return []string{""}
	}
	inputText := s.ToLower(text)
	listOfWords := s.Fields(inputText)
	return listOfWords
}

type cliCommand struct {
	name        string
	description string
	callback    func()
}

var commands map[string]cliCommand
var urls config

func init() {

	commands = map[string]cliCommand{
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
			callback:    func() { commandMap(&urls) },
		},
		"mapb": {
			name:        "mapb",
			description: "List previous pokemon locations",
			callback:    func() { commandMapB(&urls) },
		},
	}
}

//figure out how to use the cache in the program
