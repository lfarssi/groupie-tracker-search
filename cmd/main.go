package main

import (
	"groupie_tracker/database"
	"groupie_tracker/models"
	"groupie_tracker/routes"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 1 {
		os.Exit(1)
	}
	var err error
	database.Artists, err = models.GetArtists()
	for i := 0; i < len(database.Artists); i++ {
		artist := &database.Artists[i]
		if len(artist.Members) == 1 {
			artist.Type = "Artist"
		} else {
			artist.Type = "Group of " + strconv.Itoa(len(artist.Members))
		}
	}
	if err != nil {
		panic(err)
	}
	routes.Router()
}
