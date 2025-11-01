package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, additionalCommands []string) error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
