package main

import (
	"fmt"
	"math/rand"
	"time"
)

//prints available commands and a description
func commandCatch(cfg *config, pokemon string) error {
	if pokemon == "" {
		fmt.Println("Try this: catch <pokemon-name>")
		return nil
	}

	pokeResp, err := cfg.pokeapiClient.InvestigatePokemon(pokemon)
	if err != nil {
		return err
	}

	floatBaseXP := float64(pokeResp.Base_experience)

	catchrate, err := calculateCatchrate(floatBaseXP)
	if err != nil {
		return err
	}
	caught := attemptCatch(catchrate)

	fmt.Println("Throwing a ball at", pokemon)
	fmt.Println("Shake, shake...")
	if caught {
		fmt.Println(pokemon, "was caught!")
		fmt.Println("Adding to Pokedex...")
	} else {
		fmt.Println(pokemon, "escaped!")
	}

	return nil
}


func calculateCatchrate(baseXP float64) (float64, error) {
    transformedValue := 1 - ((baseXP - 20) / (608 - 20)) * (1 - 0.1)
    return transformedValue, nil
}

func attemptCatch(percentage float64) bool {
    rand.Seed(time.Now().UnixNano())
    return rand.Float64() < percentage
}