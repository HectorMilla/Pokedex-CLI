package main

import (
	"os"
	s "strings"
	"fmt"
)

func CleanInput(text string) []string{
	inputText := s.ToLower(text)
	listOfWords :=  s.Fields(inputText)
	return listOfWords
}

func commandExit(){
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
}

type cliCommand struct {
	name string
	description string
	callback func()
}

var commands = map[string]cliCommand{
	"exit":{
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	},
}