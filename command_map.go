package main

import (
	"fmt"
)

func commandMapForward(config *config) error {
	mapArea, err := config.pokeapiClient.ListLocations(config.nextLocationsUrl)
	if err != nil {
		return err
	}

	config.nextLocationsUrl = mapArea.Next
	config.prevLocationsUrl = mapArea.Previous
	for _, result := range mapArea.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandMapBack(config *config) error {
	mapArea, err := config.pokeapiClient.ListLocations(config.prevLocationsUrl)
	if err != nil {
		return err
	}

	config.nextLocationsUrl = mapArea.Next
	config.prevLocationsUrl = mapArea.Previous
	for _, result := range mapArea.Results {
		fmt.Println(result.Name)
	}

	return nil
}
