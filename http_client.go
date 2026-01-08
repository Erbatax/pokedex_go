package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func httpGet[T any](url string) (T, error) {
	res, err := http.Get(url)
	if err != nil {
		var zero T
		return zero, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		var zero T
		return zero, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		var zero T
		return zero, err
	}

	var mapArea T
	err = json.Unmarshal(body, &mapArea)
	if err != nil {
		var zero T
		return zero, err
	}

	return mapArea, nil
}
