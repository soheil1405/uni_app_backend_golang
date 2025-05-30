package models

import (
	"uni_app/database"
)

type Lessons []*Lesson

// Lesson represents a lesson
type Lesson struct {
	database.Model
	Name          string       `json:"name,omitempty"`
	Code          string       `gorm:"uniqueIndex" json:"code,omitempty"`
	UnitCount     int          `json:"unit_count,omitempty"`
	MajorID       database.PID `json:"major_id,omitempty"`
	Major         *Major       `gorm:"foreignKey:MajorID;constraint:OnDelete:CASCADE;" json:"major,omitempty"`
	DaneshKadehID database.PID `json:"daneshkadeh_id,omitempty"`
	DaneshKadeh   *DaneshKadeh `gorm:"foreignKey:DaneshKadehID;constraint:OnDelete:CASCADE;" json:"daneshkadeh,omitempty"`
	Prerequisites []Lesson     `gorm:"many2many:lesson_prerequisites;" json:"prerequisites,omitempty"`
	Description   string       `json:"description,omitempty"`
	Ratings       []Rating     `json:"ratings,omitempty" gorm:"polymorphic:Owner;polymorphicValue:lesson"`
}

type FetchLessonRequest struct {
	FetchRequest
	MajorID       database.PID `json:"major_id" query:"major_id"`
	DaneshKadehID database.PID `json:"daneshkadeh_id" query:"daneshkadeh_id"`
	Code          string       `json:"code" query:"code"`
}

func LessonAcceptIncludes() []string {
	return []string{
		"Major",
		"DaneshKadeh",
		"Prerequisites",
		"Ratings",
	}
}
