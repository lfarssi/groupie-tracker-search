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

func GetArtist(id string) (db.Artist, error) {
	var artist db.Artist
	response, err := http.Get(config.ArtistsAPIURL + "/" + id)
	if err != nil  {
		return db.Artist{} , err
	}
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&artist) ; err != nil {
		return db.Artist{}, err
	}
	return artist, nil
}