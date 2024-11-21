package routes

import (
	"groupie_tracker/controllers"
	"log"
	"net/http"
)

func Router() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", controllers.ArtistController)
	http.HandleFunc("/artist/{id}", controllers.ArtistDetailController)
	log.Println("Server running on http://localhost:8089")
	log.Fatal(http.ListenAndServe(":8089", nil))
}
