package models

import (
	"uni_app/database"
)

// DegreeLevel represents a degree level
type DegreeLevel struct {
	database.Model
	Name string `gorm:"size:255;not null" json:"name"`
}
