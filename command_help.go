package main

import (
	"fmt"
	"github.com/samassembly/pokedexcli/internal/pokecache"
)

//prints available commands and a description
func commandHelp(cfg *config, cache *pokecache.cache) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
