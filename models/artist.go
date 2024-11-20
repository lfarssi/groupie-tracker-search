package models

import (
	"encoding/json"
	"groupie_tracker/config"
	"groupie_tracker/db"
	"net/http"
)



func GetArtists() ([]db.Artist, error) {
	var artists []db.Artist
	response, err := http.Get(config.ArtistsAPIURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	
	if err := json.NewDecoder(response.Body).Decode(&artists) ; err != nil {
		return nil, err
	}
	return artists, nil
}
