package models

import (
	"errors"
	"groupie_tracker/db"
)

func ArtistById(id int) (db.Artist, error) {
   artists , err := GetArtists()
   if err!= nil {
       return db.Artist{}, err
   }
   for _, artist:= range artists {
       if artist.Id == id {
           return artist, nil
       }
    }
    return db.Artist{}, errors.New("artist not found")
}