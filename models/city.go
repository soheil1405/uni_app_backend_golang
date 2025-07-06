package models

import "uni_app/database"

type City struct {
	database.Model
	Name   string `gorm:"not null" json:"name,omitempty"`
	Unies  []Unis `json:"unies,omitempty"`
	Active bool   `json:"active"`
}

type FetchCityRequest struct {
	FetchRequest
}

func CityAcceptIncludes() []string {
	return []string{
		"Unies",
	}
}
