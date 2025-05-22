package models

import (
	"uni_app/database"
)

// DegreeLevel represents a degree level
type DegreeLevel struct {
	database.Model
	Name string `gorm:"size:255;not null" json:"name"`
	Type string `gorm:"size:255;not null" json:"type"`
}

func DegreeLevelAcceptIncludes() []string {
	return []string{}
}
