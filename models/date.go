package models

import (
	"encoding/json"
	"groupie_tracker/config"
	"net/http"
)

type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}
type Dates struct {
	Index []Date `json:"index"`
}

func GetDate() (Dates, error) {
	var dates Dates
	response, err := http.Get(config.DatesAPIURL)
	if err != nil {
		return Dates{}, err
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&dates); err != nil {
		return Dates{}, err
	}
	return dates, nil
}
