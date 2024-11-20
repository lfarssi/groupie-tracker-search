package models

import (
	"encoding/json"
	"groupie_tracker/config"
	"groupie_tracker/db"
	"net/http"
)


func GetDates(id int) ([]string, error) {
	var date db.Date
	response, err:= http.Get(config.DatesAPIURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&date) ; err != nil {
		return nil, err
	}
	return date.Index[id].Dates, nil
}