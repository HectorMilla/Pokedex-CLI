package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getPokeLocations(url string) (LocationArea, error) {
	var locArea LocationArea

	if value, ok := cache.Get(url); ok {

		locArea = value

	} else {
		res, err := http.Get(url)

		if err != nil {
			fmt.Printf("Error: %v", err)
			return LocationArea{}, fmt.Errorf("error fetching data: %w", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return LocationArea{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
		}

		if err := json.NewDecoder(res.Body).Decode(&locArea); err != nil {
			return LocationArea{}, fmt.Errorf("error decoding json: %w", err)
		}

	}
	interval, err := time.ParseDuration("5s")
	if err != nil {
		println(fmt.Errorf("unable to get time interval"))
	}
	go NewCache(interval, locArea, url)
	return locArea, nil
}
