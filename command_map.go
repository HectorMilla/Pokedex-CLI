package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	locationArea, _ := cfg.pokeapiClient.GetPokeLocations(&cfg.next)
	cfg.next = locationArea.Next
	cfg.previous = fmt.Sprint(locationArea.Previous)

	for _, area := range locationArea.Results {
		fmt.Println(area.Name)
	}
	return nil
}
