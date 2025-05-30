package models

import "uni_app/database"

type Terms []*Term

type Term struct {
	database.Model
	Name        string `json:"name,omitempty"`
	StartDate   string `json:"start_date,omitempty"`
	EndDate     string `json:"end_date,omitempty"`
	IsActive    bool   `gorm:"default:false" json:"is_active,omitempty"`
	Description string `json:"description,omitempty"`
}

type FetchTermRequest struct {
	FetchRequest
	IsActive bool `json:"is_active" query:"is_active"`
}

func TermAcceptIncludes() []string {
	return []string{}
}
