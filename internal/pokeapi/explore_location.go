package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ExploreLocation -
func (c *Client) ExploreLocation(location string) (RespDeepLocation, error) {
	url := baseURL + "/location-area/" + location

	if val, ok := c.cache.Get(url); ok {
		exploreResp := RespDeepLocation{}

		err := json.Unmarshal(val, &exploreResp)
		if err != nil {
			return RespDeepLocation{}, err
		}

		return exploreResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDeepLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDeepLocation{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDeepLocation{}, err
	}

	exploreResp := RespDeepLocation{}
	err = json.Unmarshal(dat, &exploreResp)
	if err != nil {
		return RespDeepLocation{}, err
	}

	c.cache.Add(url, dat)
	return exploreResp, nil
}
