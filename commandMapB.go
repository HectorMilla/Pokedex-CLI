package main

import (
	"fmt"
	"strings"
)

func commandMapB(urls *config) {
	if urls.previous == "(0x0,0x0)" || !strings.Contains(urls.previous, "http") {
		fmt.Println("Use command map to get a list of locations")
		return
	}
	locationArea, _ := getPokeLocations(urls.previous)
	for _, area := range locationArea.Results {
		fmt.Println(area.Name)
	}

}
