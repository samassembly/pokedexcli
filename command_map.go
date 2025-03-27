package main

import (
	"errors"
	"fmt"
	"github.com/samassembly/pokedexcli/internal/pokecache"
)

func commandMapf(cfg *config, cache *pokecache.cache) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	cacheVal, exists := cache.Get(cfg.nextLocationsURL)
	if exists {
		fmt.Println("cache is present")
		body = cacheVal
	} else {
		fmt.Println("cache is not present")
		for _, loc := range locationsResp.Results {
			fmt.Println(loc.Name)
		}
		cache.Add(locationsResp.Next, locationsResp.Results)
	}

	return nil
}

func commandMapb(cfg *config, cache *pokecache.cache) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}