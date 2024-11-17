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

func GetArtist(id string) (Artist, error) {
	var artist Artist
	response, err := http.Get(config.ArtistsAPIURL + "/" + id)
	if err != nil  {
		return Artist{} , err
	}
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&artist) ; err != nil {
		return Artist{}, err
	}
	return artist, nil
}