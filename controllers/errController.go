package controllers

import (
	"groupie_tracker/models"
	"html/template"
	"net/http"
)

func ErrorController(w http.ResponseWriter, r *http.Request, statusCode int) {
	tmp, err := template.ParseFiles("views/error.html")
	if err != nil {
		http.Error(w,"500 | Internal Server Error", http.StatusInternalServerError)
		return
	}
	errorType := models.Error{
		StatusCode: statusCode,
		Message: http.StatusText(statusCode),
	}
	w.WriteHeader(statusCode)
	if err:= tmp.Execute(w, errorType) ; err != nil {	
		http.Error(w, "500 | Internal Server Error", http.StatusInternalServerError)
		return
	}
}
