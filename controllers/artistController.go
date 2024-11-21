package controllers

import (
	"groupie_tracker/models"
	"html/template"
	"net/http"
)

func ArtistController(w http.ResponseWriter, r *http.Request) {
	artist, err := models.GetArtists()
	if err != nil {
		ErrorController(w, r, http.StatusInternalServerError)
		return
	}
	res, err1 := template.ParseFiles("views/index.html")
	if err1 != nil {
		ErrorController(w, r, http.StatusInternalServerError)
		return
	}

	if err := res.Execute(w, artist); err != nil {
		ErrorController(w, r, http.StatusInternalServerError)
		return
	}
}
