package main

import "fmt"

func commandMap(urls *config) {
	if urls.next == "" {
		urls.next = "https://pokeapi.co/api/v2/location-area"
		urls.offset = 20
	}
	locationArea, _ := getPokeLocations(urls.next)
	urls.next = locationArea.Next
	urls.previous = fmt.Sprint(locationArea.Previous)

	for _, area := range locationArea.Results {
		fmt.Println(area.Name)
	}
}
