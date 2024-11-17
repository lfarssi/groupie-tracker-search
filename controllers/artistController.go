package controllers

import (
	"groupie_tracker/models"
	"net/http"
	"html/template"
)

func ArtistController(w http.ResponseWriter, r *http.Request) {
	artist, err := models.GetArtists()
	if err != nil {
		http.Error(w,  "Failed to parse template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl, err1  := template.ParseFiles("views/index.html")
	if err1 != nil {
		http.Error(w, "Failed to parse template: "+err1.Error(), http.StatusInternalServerError)
		return
	}
	
	if err := tmpl.Execute(w, artist) ; err!= nil {
        http.Error(w,  "Failed to render template: "+err.Error(), http.StatusInternalServerError)
        return
    }
}