package models

import (
	"uni_app/database"
)

type FetchDepartmentRequest struct {
	FetchRequest
	FacultyID database.PID `json:"faculty_id,omitempty"`
	Name      string       `json:"name,omitempty"`
}

type Department struct {
	database.Model
	FacultyID   database.PID `gorm:"not null" json:"faculty_id"`
	Faculty     *Faculty     `json:"faculty,omitempty"`
	Name        string       `gorm:"not null" json:"name"`
	Description string       `json:"description,omitempty"`
	HeadID      database.PID `json:"head_id,omitempty"`
	Head        *Teacher     `json:"head,omitempty"`
	Teachers    []*Teacher   `json:"teachers,omitempty" gorm:"many2many:department_teachers;"`
	Majors      []*Major     `json:"majors,omitempty"`
}
