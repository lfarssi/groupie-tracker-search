package controllers

import (


	"groupie_tracker/models"
	"html/template"
	"net/http"
	"strconv"
)

// func fetchData(url string, target interface{}) error {
// 	 response, err := http.Get(url)
// 	 if err != nil {
// 		return err 
// 	 }
// 	 defer response.Body.Close()
// 	 return json.NewDecoder(response.Body).Decode(target)
// }

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
	// var locations db.Location
	// if err := fetchData(config.LocationsAPIURL, &locations) ; err != nil {
	// 	http.Error(w, "Failed to fetch location: "+err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// var dates db.Date
	// if err := fetchData(config.DatesAPIURL, &dates) ; err != nil {
	// 	http.Error(w, "Failed to fetch location: "+err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// var relations db.Relation
	// if err := fetchData(config.RelationAPIURL, &relations) ; err != nil {
	// 	http.Error(w, "Failed to fetch location: "+err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	res, err := template.ParseFiles("views/detail.html")
	if err != nil {
		http.Error(w, "Failed to parse template: "+err.Error(), http.StatusInternalServerError)
		return
	}
		// Extract datesLocations from relations

		// Combine all data into a single struct
		// artistDetail := db.Artist{
		// 	Id:             artist.Id,
		// 	Name:           artist.Name,
		// 	Image:          artist.Image,
		// 	Members:        artist.Members,
		// 	CreationDate:   artist.CreationDate,
		// 	FirstAlbum:     artist.FirstAlbum,
		// 	Locations: locations.Index[id].Location[id],
		// 	ConcertDates: dates.Index[id].Dates[id],
		// 	Relations: relations.Index[id].DatesLocations,
		// }
	
		res.Execute(w, artist)
}
