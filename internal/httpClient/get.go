package httpclient

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// ListLocations -
func Get[T any](c *Client, url *string) (T, error) {
	requestUrl := c.baseUrl
	if url != nil {
		if strings.HasPrefix(*url, c.baseUrl) == false {
			requestUrl += *url
		} else {
			requestUrl = *url
		}
	}

	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		var zero T
		return zero, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		var zero T
		return zero, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		var zero T
		return zero, err
	}

	var locationsResp T
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		var zero T
		return zero, err
	}

	return locationsResp, nil
}
