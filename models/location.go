package models

import (
	"encoding/json"
	"groupie_tracker/config"
	"net/http"
)


type Location struct {
	Id int `json:"id"`
	Locations []string `json:"locations"`

}

func GetLocations() ([]Location, error) {
	var locations []Location
	response, err:= http.Get(config.LocationsAPIURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&locations) ; err != nil {
		return nil, err
	}
	return locations, nil
}