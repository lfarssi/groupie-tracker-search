package models

import (
	"encoding/json"
	"groupie_tracker/config"
	"net/http"
)


type Artist struct {
	Id          int      `json:"id"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	Image        string   `json:"image"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`     
	ConcertDates string   `json:"concertDates"`  
	Relations    string  `json:"relations"`
	Relation map[string][]string 
}

func GetArtists() ([]Artist, error) {
	var artists []Artist
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
