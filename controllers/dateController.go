package controllers

import (
	"fmt"
	"groupie_tracker/models"
	"html/template"
	"net/http"
	"strconv"
)

func DateController(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Date ID", http.StatusBadRequest)
		return
	}

	dates, err := models.GetDates(id)
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
