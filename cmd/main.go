package main

import (
	"groupie_tracker/routes"
	"groupie_tracker/database"
	"groupie_tracker/models"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		os.Exit(1)
	}
	var err error
	database.Artists, err = models.GetArtists()
	if err!= nil {
        panic(err)
    }
		routes.Router()
}