package controllers

import (
	"groupie_tracker/models"
	"html/template"
	"net/http"
	"strconv"
)

func ArtistDetailController(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}
	artist, err := models.ArtistById(id)
	if err != nil {
		http.Error(w, "Failed to fetch artist: "+err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := template.ParseFiles("views/detail.html")
	if err != nil {
		http.Error(w, "Failed to parse template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	res.Execute(w, artist)
}
