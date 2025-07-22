package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	if _, ok := cfg.pokemon[pokemon.Name]; ok {
		fmt.Printf("You already have %s...\n", pokemon.Name)
		return nil
	}
	baseExp := pokemon.BaseExperience

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	randRes := rand.Intn(baseExp)
	if randRes > baseExp/2 {
		cfg.pokemon[pokemon.Name] = pokemon
		fmt.Printf("%s was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}
