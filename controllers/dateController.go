package controllers

import (
	"fmt"
	"groupie_tracker/models"
	"html/template"
	"net/http"

)

func DateController(w http.ResponseWriter, r *http.Request) {

	dates, err := models.GetDates()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := template.ParseFiles("views/detail.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Print(dates)

	if err = res.Execute(w, dates); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
