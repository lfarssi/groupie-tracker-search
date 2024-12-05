package models

import (
	"encoding/json"
	"groupie_tracker/config"
	"net/http"
)

type Location struct {
	Id       int      `json:"id"`
	Location []string `json:"locations"`
}
type Locations struct {
	Index []Location `json:"index"`
}

func GetLocation() (Locations, error) {
	var locations Locations
	response, err := http.Get(config.LocationsAPIURL)
	if err != nil {
		return Locations{}, err
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&locations); err != nil {
		return Locations{}, err
	}
	return locations, nil
}
