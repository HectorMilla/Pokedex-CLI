package main

import (
	"fmt"
	"strings"
)

func commandMap(cfg *config, additionalCommands []string) error {
	locationArea, _ := cfg.pokeapiClient.GetPokeLocations(&cfg.next)
	cfg.next = locationArea.Next
	cfg.previous = fmt.Sprint(locationArea.Previous)

	for _, area := range locationArea.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapB(cfg *config, additionalCommands []string) error {
	if cfg.previous == "(0x0,0x0)" || !strings.Contains(cfg.previous, "http") {
		return fmt.Errorf("use command map to get a list of locations")
	}
	locationArea, _ := cfg.pokeapiClient.GetPokeLocations(&cfg.previous)
	cfg.next = locationArea.Next
	cfg.previous = fmt.Sprint(locationArea.Previous)
	for _, area := range locationArea.Results {
		fmt.Println(area.Name)
	}
	return nil
}
