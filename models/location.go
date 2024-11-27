package models

import (
	"encoding/json"
	"groupie_tracker/config"
	"net/http"
)


type Location struct {
	Id       int    `json:"id"`
	Location []string `json:"locations"`
}

func GetLocation() ([]Location, error) {
	var location []Location
	response, err := http.Get(config.LocationsAPIURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&location); err != nil {
		return nil, err
	}
	return location, nil
}
