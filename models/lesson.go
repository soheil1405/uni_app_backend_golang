package models

import (
	"uni_app/database"
)

// Lesson represents a lesson
type Lesson struct {
	database.Model
	Name        string `gorm:"size:255;not null" json:"name"`
	Description string `gorm:"size:500" json:"description"`
}
