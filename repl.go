package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	config := &PaginationConfig{}
	reader := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("Pokedex > ")
        reader.Scan()

        words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config)
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
	callback    func(*PaginationConfig) error
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
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Displays previous page of location areas",
		callback:    commandMapB,
	},
	}
}

type PaginationConfig struct {
	Next string
	Previous *string
}