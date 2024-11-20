package models

import (
	"encoding/json"
	"groupie_tracker/config"
	"groupie_tracker/db"
	"net/http"
)


func GetLocations() ([]string, error) {
	var locations db.Location
	response, err:= http.Get(config.LocationsAPIURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&locations) ; err != nil {
		return nil, err
	}
	return locations.Location, nil
}