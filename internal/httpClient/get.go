package httpClient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ListLocations -
func Get[T any](c *Client, url *string) (T, error) {
	requestUrl := c.constructUrl(url)

	dat, ok := c.cache.Get(requestUrl)
	if !ok {
		req, err := http.NewRequest("GET", requestUrl, nil)
		if err != nil {
			var zero T
			return zero,
				fmt.Errorf("Error creating request: %v\n", err)
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			var zero T
			return zero,
				fmt.Errorf("Error making HTTP request: %v\n", err)
		}
		defer resp.Body.Close()

		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			var zero T
			return zero,
				fmt.Errorf("Error reading response body: %v\n", err)
		}

		if resp.StatusCode > 299 {
			var zero T
			return zero, fmt.Errorf("Response failed with status code: %d\n", resp.StatusCode)
		}

		c.cache.Add(requestUrl, dat)
	}

	var locationsResp T
	err := json.Unmarshal(dat, &locationsResp)
	if err != nil {
		var zero T
		return zero,
			fmt.Errorf("Error unmarshalling JSON: %v\n", err)
	}

	return locationsResp, nil
}
