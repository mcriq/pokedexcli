package main

import (
	"fmt"
)

func commandMap(config *PaginationConfig) error {
	urlToUse := "https://pokeapi.co/api/v2/location-area/"
	if config.Next != "" {
		urlToUse = config.Next
	}

	locationData, err := getLocationArea(urlToUse)
	if err != nil {
		return fmt.Errorf("unable to fetch location data: %v", err)
	}

	// Print the locations
	for _, area := range locationData.Results {
		fmt.Println(area.Name)
	}

	// Update the config for next time
	config.Next = locationData.Next
	config.Previous = locationData.Previous
	return nil
}