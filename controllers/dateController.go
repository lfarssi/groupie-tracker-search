package controllers

import (
	"groupie_tracker/database"
	"html/template"
	"net/http"
)

func DateController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorController(w, r, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/dates" {
		ErrorController(w, r, http.StatusNotFound)
		return
	}

	dates := database.Dates
	res, err1 := template.ParseFiles("views/dates.html")
	if err1 != nil {
		ErrorController(w, r, http.StatusInternalServerError)
		return
	}

	if err := res.Execute(w, dates); err != nil {
		ErrorController(w, r, http.StatusInternalServerError)
		return
	}
}
