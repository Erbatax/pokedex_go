package main

import (
	"fmt"
)

func commandMapForward(config *config, _ ...string) error {
	mapArea, err := config.pokeapiClient.GetLocationAreaList(config.nextLocationsUrl)
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

func commandMapBack(config *config, _ ...string) error {
	mapArea, err := config.pokeapiClient.GetLocationAreaList(config.prevLocationsUrl)
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
