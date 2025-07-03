package models

import (
	"uni_app/database"
)

type FetchFacultyRequest struct {
	FetchRequest
	UniversityID database.PID `json:"university_id,omitempty"`
	Name         string       `json:"name,omitempty"`
}

type Faculty struct {
	database.Model
	UniversityID database.PID  `gorm:"not null" json:"university_id"`
	University   *Uni   `json:"university,omitempty"`
	Name         string        `gorm:"not null" json:"name"`
	Description  string        `json:"description,omitempty"`
	Floors       []*Floor      `json:"floors,omitempty"`
	Departments  []*Department `json:"departments,omitempty"`
	Students     []*User       `json:"students,omitempty" gorm:"many2many:faculty_students;"`
	Teachers     []*Teacher    `json:"teachers,omitempty" gorm:"many2many:faculty_teachers;"`
	Staff        []*User       `json:"staff,omitempty" gorm:"many2many:faculty_staff;"`
}
