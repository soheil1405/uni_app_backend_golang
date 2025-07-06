package models

import (
	"uni_app/database"
)

type Floor struct {
	database.Model
	Active      bool         `json:"active"`
	FacultyID   database.PID `gorm:"not null" json:"faculty_id"`
	Faculty     *Faculty     `json:"faculty,omitempty"`
	Description string       `json:"description,omitempty"`
	Rooms       []*Room      `json:"rooms,omitempty"`
}
