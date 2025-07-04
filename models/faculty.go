package models

import (
	"uni_app/database"
)

type FacultyType string

const (
	FacultyTypeFanni    FacultyType = "fanni"
	FacultyTypeEnsani   FacultyType = "ensani"
	FacultyTypeMemari   FacultyType = "memari"
	FacultyTypeModiriat FacultyType = "modiriat"
	FacultyTypeHonar    FacultyType = "honar"
)

type FetchFacultyRequest struct {
	FetchRequest
	UniID database.PID `gorm:"not null" json:"uni_id,omitempty"`
	Name  string       `json:"name,omitempty"`
}

type Faculty struct {
	database.Model
	Name        string        `gorm:"not null" json:"name,omitempty"`
	Description string        `json:"description,omitempty"`
	UniID       database.PID  `gorm:"not null" json:"uni_id,omitempty"`
	Uni         *Uni          `json:"university,omitempty"`
	FacultyType FacultyType   `json:"faculty_type,omitempty"`
	Address     *Address      `gorm:"polymorphic:Owner;auto_preload:false;" json:"address,omitempty"`
	ContactWays []*ContactWay `gorm:"polymorphic:Owner;auto_preload:false;" json:"contact_ways,omitempty"`
	Courses     []*Course     `json:"courses,omitempty" gorm:"many2many:faculty_courses;"`
	Floors      []*Floor      `json:"floors,omitempty"`
	Students    Student       `json:"students,omitempty"`
	Users       Users         `json:"staff,omitempty" gorm:"many2many:faculty_staff;"`
	Ratings     []*Rating     `json:"ratings,omitempty" gorm:"polymorphic:Owner;polymorphicValue:faculties"`
	Majors      []*Major      `json:"majors,omitempty"`
}
