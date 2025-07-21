package main

import (
	"fmt"

	"github.com/MS-32154/pokeapi"
)

func commandMap(config *Config) error {
	url := config.Next
	areas, next, previous, err := pokeapi.GetLocationAreas(url)
	if err != nil {
		return err
	}
	for _, area := range areas {
		fmt.Printf("%s\n", area.Name)
	}

	config.Next = next

	config.Previous = previous

	return nil
}

func commandMapb(config *Config) error {
	url := config.Previous
	areas, next, previous, err := pokeapi.GetLocationAreas(url)
	if err != nil {
		return err
	}
	for _, area := range areas {
		fmt.Printf("%s\n", area.Name)
	}

	config.Next = next

	config.Previous = previous

	return nil
}
