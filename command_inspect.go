package main

import (
	"fmt"
)

//prints available commands and a description
func commandInspect(cfg *config, pokemon string) error {
	if _, exists := cfg.caughtPokemon[pokemon]; exists {
		fmt.Printf("Name: %s\n", cfg.caughtPokemon[pokemon].Name)
        fmt.Printf("Height: %d\n", cfg.caughtPokemon[pokemon].Height)
        fmt.Printf("Weight: %d\n", cfg.caughtPokemon[pokemon].Weight)
		for _, Type := range cfg.caughtPokemon[pokemon].Types {
			fmt.Printf(" - %s\n", Type.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}
