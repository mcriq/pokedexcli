package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getLocationArea(url string) (LocationArea, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationArea{}, fmt.Errorf("unable to get location areas: %v", err)
	}
	defer res.Body.Close()

	area := LocationArea{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&area); err != nil {
		return LocationArea{}, fmt.Errorf("error decoding response body: %v", err)
	}
	return area, nil
}