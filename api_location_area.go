package main

import (
	"strings"
)

type MapArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getLocationArea(url string) (MapArea, error) {
	baseUrl := "https://pokeapi.co/api/v2/location-area/"
	if strings.HasPrefix(url, baseUrl) == false {
		url = baseUrl
	}

	return httpGet[MapArea](url)
}
