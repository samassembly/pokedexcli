package main

import (
	"fmt"
	"os"
	"github.com/samassembly/pokedexcli/internal/pokecache"
)

//exits the repl loop and closes program
func commandExit(cfg *config, cache *pokecache.cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
