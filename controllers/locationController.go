package controllers

import (
	"groupie_tracker/models"
	"html/template"
	"net/http"
)

func LocationController(w http.ResponseWriter, r *http.Request){
	locations, err := models.GetLocations()
	if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	res, err := template.ParseFiles("views/locations.html")
	if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	err = res.Execute(w, locations)
	if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}