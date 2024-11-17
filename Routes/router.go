package routes

import (
	"groupie_tracker/controllers"
	"log"
	"net/http"
)

func Router() {
	http.HandleFunc("/", controllers.ArtistController)
	http.HandleFunc("/artist/{id}", controllers.ArtistDetailController)
	http.HandleFunc("/locations", controllers.LocationController)
	log.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}