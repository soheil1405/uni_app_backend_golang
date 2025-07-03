package models

import (
	"uni_app/database"
)

type Floor struct {
	database.Model
	FacultyID   database.PID `gorm:"not null" json:"faculty_id"`
	Faculty     *Faculty     `json:"faculty,omitempty"`
	Number      int          `gorm:"not null" json:"number"`
	Description string       `json:"description,omitempty"`
	Rooms       []*Room      `json:"rooms,omitempty"`
}
