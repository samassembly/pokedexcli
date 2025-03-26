package main

import (
	"fmt"
	"errors"
)

//get the next 20 locations from pokeapi
func commandMapf(cfg *config) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		log.Fatal(err)
	}

	//updates the structs next and previous location calls
	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	//print locations
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

//get the previous 20 locations from pokeapi
func commandMapb(cfg *config) error {
	//check if first
	if cfg.prevLocationsURL == nil {
		return errors.New("This is the first page of locations")
	}
	//get batch of locations
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}
	//updates the structs next and previous location calls
	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.previous

	//print locations
	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil

}