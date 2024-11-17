package controllers

import (
	"encoding/json"
	"groupie_tracker/models"
	"html/template"
	"net/http"
	"strconv"
)

func fetchData(url string, target interface{}) error {
	 response, err := http.Get(url)
	 if err != nil {
		return err 
	 }
	 defer response.Body.Close()
	 return json.NewDecoder(response.Body).Decode(target)
}

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
	var locations []string
	if err := fetchData(artist.Locations, &locations) ; err != nil {
		http.Error(w, "Failed to fetch location: "+err.Error(), http.StatusInternalServerError)
		return
	}
	var dates []string
	if err := fetchData(artist.ConcertDates, &dates) ; err != nil {
		http.Error(w, "Failed to fetch location: "+err.Error(), http.StatusInternalServerError)
		return
	}
	var relations []string
	if err := fetchData(artist.ConcertDates, &relations) ; err != nil {
		http.Error(w, "Failed to fetch location: "+err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := template.ParseFiles("views/detail.html")
	if err != nil {
		http.Error(w, "Failed to parse template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	res.Execute(w, artist)
}
