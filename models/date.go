package models

import (
	"encoding/json"
	"groupie_tracker/config"
	"net/http"
)

type Date struct {
	Id int `json:"id"`
	Dates []string `json:"dates"`
}

func GetDate() ([]Date, error) {
	var date []Date
	response, err := http.Get(config.DatesAPIURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&date); err != nil {
		return nil, err
	}
	return date, nil
}
