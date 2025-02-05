package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	locationName := ""

	if len(args) == 0 {
		return errors.New("no location entered")
	}

	locationName = args[0]

	encountersResp, err := cfg.pokeapiClient.ListEncounters(locationName)
	if err != nil {
		return err
	}

	// check if the PokemonEncounters slice is empty and then loop over it if not empty
	if len(encountersResp.PokemonEncounters) == 0 {
		return fmt.Errorf("no pokemon found in: %v", encountersResp.Name)
	}

	fmt.Printf("Exploring %s...\n", locationName)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range encountersResp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}