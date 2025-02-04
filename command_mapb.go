package main

import (
	"fmt"
)

func commandMapB(config *PaginationConfig) error {
	if config.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	urlToUse := *config.Previous

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