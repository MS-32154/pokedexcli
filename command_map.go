package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	if cfg.nextLocationsURL != nil {
		v, exists := cfg.cache.Get(*cfg.nextLocationsURL)
		if exists {
			for _, loc := range v {
				fmt.Println(loc)
			}
			return nil
		}
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL != nil {
		v, exists := cfg.cache.Get(*cfg.prevLocationsURL)
		if exists {
			for _, loc := range v {
				fmt.Println(loc)
			}
			return nil
		}
	}

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
