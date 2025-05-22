package models

import "uni_app/database"

type City struct {
	database.Model
	Name  string `gorm:"not null" json:"name,omitempty"`
	Unies []Unis `json:"unies,omitempty"`
}

type FetchCityRequest struct {
	FetchRequest
}

func CityAcceptIncludes() []string {
	return []string{
		"Unies",
	}
}
