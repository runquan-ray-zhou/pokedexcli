package main

import (
	"fmt"
	"log"
)

func callbackMap(cfg *config) error {
	// pokeapiClient := pokeapi.NewClient()

	resp, err := cfg.pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}
