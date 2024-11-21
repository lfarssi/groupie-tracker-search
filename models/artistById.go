package models

import (
	"errors"
)

func ArtistById(id int) (Artist, error) {
   artists , err := GetArtists()
   if err!= nil {
       return Artist{}, err
   }
   for _, artist:= range artists {
       if artist.Id == id {
           return artist, nil
       }
    }
    return Artist{}, errors.New("artist not found")
}