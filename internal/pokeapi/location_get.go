package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListEncounters(location string) (Location, error) {
	url := baseURL + "/location-area/" + location

	if cached, ok := c.cache.Get(url); ok {
		var encounters Location
		err := json.Unmarshal(cached, &encounters)
		if err != nil {
			return Location{}, err
		}
		return encounters, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, dat)

	encounterResp := Location{}
	err = json.Unmarshal(dat, &encounterResp)
	if err != nil {
		return Location{}, err
	}

	return encounterResp, nil
}