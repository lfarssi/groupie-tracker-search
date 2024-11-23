package controllers

import (
	"groupie_tracker/database"
	"html/template"
	"net/http"
)

func RelationController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorController(w, r, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/relations" {
		ErrorController(w, r, http.StatusNotFound)
		return
	}

	relations := database.Relations
	res, err1 := template.ParseFiles("views/relations.html")
	if err1 != nil {
		ErrorController(w, r, http.StatusInternalServerError)
		return
	}

	if err := res.Execute(w, relations); err != nil {
		ErrorController(w, r, http.StatusInternalServerError)
		return
	}
}
