package models

import "uni_app/database"

type MajorType string

// رشته تحصیلی
type Major struct {
	database.Model
	DegreeLevel DegreeLevel `json:"degree_level,omitempty"`
	Name        string      `gorm:"not null" json:"name,omitempty"`
	Code        string      `gorm:"unique;not null" json:"code,omitempty"`
	Description string      `json:"description,omitempty"`
	Students    []Student   `gorm:"foreignKey:MajorID;constraint:OnDelete:CASCADE;" json:"students,omitempty"`
}

type FetchMajorRequest struct {
	FetchRequest
	DegreeLevel DegreeLevel `json:"-,omitempty" query:"degree_level"`
}

func MajorAcceptIncludes() []string {
	return []string{
		"Students",
	}
}
