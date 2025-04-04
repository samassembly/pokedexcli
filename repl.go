package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/samassembly/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon map[string]pokeapi.PokemonInfo
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
		input := ""
		if len(words) > 1 {
			input = words[1]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, input)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
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
	callback    func(*config, string) error
}

func commandPokedex(cfg *config, inp string) error {
	fmt.Println("Your Pokedex:")
	for _, Pok := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", Pok.Name)
	}
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Get pokemon in given location",
			callback:    commandExplore,
		},
		"catch": {
			name: 		 "catch",
			description: "Attempt to catch a given pokemon",
			callback: 	 commandCatch,
		},
		"inspect": {
			name: 		 "inspect",
			description: "Inspect details for a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name: 		 "pokedex",
			description: "Prints list of all caught pokemon",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
