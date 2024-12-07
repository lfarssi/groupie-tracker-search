package controllers

import (
	"groupie_tracker/models"
	"strconv"
	"strings"
)

func Search(artists []models.Artist, search string) (any, error) {
	search = strings.ToLower(search)
	searched := make([]models.Artist, 0)
	uniqueArtists := make(map[int]struct{}) // To track unique artists by their Id

	for _, artist := range artists {
		// Skip if artist has already been added
		if _, exists := uniqueArtists[artist.Id]; exists {
			continue
		}

		// Check if artist matches search criteria
		if strings.Contains(strings.ToLower(artist.Name), search) ||
			strings.Contains(artist.FirstAlbum, search) ||
			strings.Contains(strconv.Itoa(artist.CreationDate), search) {
			searched = append(searched, artist)
			uniqueArtists[artist.Id] = struct{}{}
			continue
		}

		// Check if any member matches the search criteria
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), search) {
				searched = append(searched, artist)
				uniqueArtists[artist.Id] = struct{}{}
				break
			}
		}

		// Check if any concert date matches the search criteria
		for _, date := range artist.ConcertDatesr {
			if strings.Contains(strings.ToLower(date), search) {
				searched = append(searched, artist)
				uniqueArtists[artist.Id] = struct{}{}
				break
			}
		}

		// Check if any location matches the search criteria
		for _, location := range artist.Locationsr {
			if strings.Contains(strings.ToLower(location), search) {
				searched = append(searched, artist)
				uniqueArtists[artist.Id] = struct{}{}
				break
			}
		}
	}
	
	return searched, nil
}
