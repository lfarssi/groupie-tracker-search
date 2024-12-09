package main

import (
	"groupie_tracker/database"
	"groupie_tracker/models"
	"groupie_tracker/routes"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 1 {
		os.Exit(1)
	}
	database.Artists, _ = models.GetArtists()
	database.Locations, _ = models.GetLocation()
	database.Dates, _ = models.GetDate()
	for i := 0; i < len(database.Artists); i++ {
		artist := &database.Artists[i]
		if len(artist.Members) == 1 {
			artist.Type = "Artist"
		} else {
			artist.Type = "Group of " + strconv.Itoa(len(artist.Members))
		}
	}
	routes.Router()
}
