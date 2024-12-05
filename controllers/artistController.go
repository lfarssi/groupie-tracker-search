package controllers

import (
	"groupie_tracker/database"
	"html/template"
	"net/http"
)

func ArtistController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorController(w, r, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		ErrorController(w, r, http.StatusNotFound)
		return
	}
	artists := database.Artists
	locations := database.Locations
	dates := database.Dates
	for i, _ := range artists {
			artists[i].Locationsr = locations.Index[i].Location		
			artists[i].ConcertDatesr = dates.Index[i].Dates
		
	}
	search := r.URL.Query().Get("search")
	if len(search) > 0 {
		data, _ := Search(artists, search)
		res, err1 := template.ParseFiles("views/index.html")
		if err1 != nil {
			ErrorController(w, r, http.StatusInternalServerError)
			return
		}
		res.Execute(w, data)
	} else {
		res, err1 := template.ParseFiles("views/index.html")
		if err1 != nil {
			ErrorController(w, r, http.StatusInternalServerError)
			return
		}
	
		if err := res.Execute(w, artists); err != nil {
			ErrorController(w, r, http.StatusInternalServerError)
			return
		}
	}


}
