package models

import (
	"encoding/json"
	"groupie_tracker/config"
	"net/http"
)


type Relation struct {
	Id            int       `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`

} 

func GetRelation() ([]Relation, error) {
	var relation []Relation
	response, err := http.Get(config.RelationAPIURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&relation); err != nil {
		return nil, err
	}
	return relation, nil
}
