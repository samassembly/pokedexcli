package main

import (
	//"errors"
	"fmt"
)
//test and production signatures
//func commandExplore(cfg *config, location string) error {
func commandExplore(cfg *config, location string) error {
	//test line
	//location := "pastoria-city-area"
	exploreResp, err := cfg.pokeapiClient.ExploreLocation(location)
	if err != nil {
		return err
	}


	fmt.Printf("Exploring %s...", location)
	fmt.Println()
	fmt.Println("Found Pokemon:")

	for _, pok := range exploreResp.PokemonEncounters {
		fmt.Println(" - ", pok.Pokemon.Name)
	}
	return nil
}