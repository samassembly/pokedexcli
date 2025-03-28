package main

import (
	"fmt"
	"os"
)

//exits the repl loop and closes program
func commandExit(cfg *config, inp string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
