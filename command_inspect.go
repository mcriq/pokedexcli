package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	pokemonName := ""

	if len(args) < 1 {
		return errors.New("no pokemon name entered")
	}

	pokemonName = args[0]

	// check if pokemon is in pokedex
	if pokemon, isCaught := cfg.caughtPokemon[pokemonName]; isCaught {
		fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", pokemon.Name, pokemon.Height, pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, pokeType := range pokemon.Types {
			fmt.Printf("  - %s\n", pokeType.Type.Name)
		}
		return nil
	}
	fmt.Println("You have not caught that pokemon")
	return nil
}