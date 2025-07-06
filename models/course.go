package models

import (
	"uni_app/database"
)

type FetchCourseRequest struct {
	FetchRequest
	Name string `json:"name,omitempty"`
}

type Course struct {
	database.Model
	Active          bool              `json:"active"`
	Name            string            `gorm:"not null" json:"name,omitempty"`
	UnitCount       int               `json:"unit_count,omitempty"`
	UniID           database.PID      `json:"uni_id,omitempty"`
	Uni             Uni               `json:"uni,omitempty"`
	MajorID         database.PID      `json:"major_id,omitempty"`
	Major           *Major            `gorm:"foreignKey:MajorID;constraint:OnDelete:CASCADE;" json:"major,omitempty"`
	RecommendedTerm int               `json:"recommended_term,omitempty"`
	IsOptional      bool              `gorm:"not null" json:"is_optional,omitempty"`
	IsTechnical     bool              `gorm:"not null" json:"is_technical,omitempty"`
	Prerequisites   []Course          `gorm:"many2many:course_prerequisites;" json:"prerequisites,omitempty"`
	Description     string            `json:"description,omitempty"`
	CourseInstances []*CourseInstance `json:"course_instances,omitempty"`
}

type CoursePrerequisite struct {
	database.Model
	CourseID       database.PID `gorm:"not null" json:"course_id"`
	Course         *Course      `json:"course,omitempty"`
	PrerequisiteID database.PID `gorm:"not null" json:"prerequisite_id"`
	Prerequisite   *Course      `json:"prerequisite,omitempty"`
}
