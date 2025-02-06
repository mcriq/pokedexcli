package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	pokemonName := ""

	if len(args) == 0 {
		return errors.New("no pokemon name entered")
	}

	pokemonName = args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	
	res := rand.Intn(pokemon.BaseExperience)
	
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)


	if res > 40 {
		fmt.Printf("%s has escaped!\n", pokemonName)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonName)
	cfg.caughtPokemon[pokemonName] = pokemon
	return nil
}

