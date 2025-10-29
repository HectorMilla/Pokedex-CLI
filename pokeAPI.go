package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	res, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error: %v", err)
		return LocationArea{}, fmt.Errorf("error fetching data: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return LocationArea{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	var locArea LocationArea
	if err := json.NewDecoder(res.Body).Decode(&locArea); err != nil {
		return LocationArea{}, fmt.Errorf("error decoding json: %w", err)
	}

	return locArea, nil
}
