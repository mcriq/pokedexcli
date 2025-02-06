package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mcriq/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	caughtPokemon map[string]pokeapi.Pokemon
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("Pokedex > ")
        reader.Scan()

        words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unkown command")
			continue
		}
    }
}

func cleanInput(text string) []string {
    output := strings.ToLower(text)
	words := strings.Fields(output)
    return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Displays location areas",
		callback:    commandMapf,
	},
	"mapb": {
		name:        "mapb",
		description: "Displays previous page of location areas",
		callback:    commandMapb,
	},
	"explore": {
		name:        "explore <location_name>",
		description: "Let's you explore a location and displays the pokemon in that area",
		callback:    commandExplore,
	},
	"catch": {
		name:        "catch <pokemon_name>",
		description: "Let's you attempt to catch a pokemon",
		callback:    commandCatch,
	},
	"inspect": {
		name:        "inspect <pokemon_name>",
		description: "Displays pokedex information on a pokemon if they have been caught",
		callback:    commandInspect,
	},
	"pokedex": {
		name:        "pokedex",
		description: "Displays all pokemon in your pokedex",
		callback:    commandPokedex,
	},
	}
}

