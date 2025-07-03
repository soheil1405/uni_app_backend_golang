package models

import (
	"uni_app/database"
)

type FetchCourseRequest struct {
	FetchRequest
	Name    string `json:"name,omitempty"`
	Code    string `json:"code,omitempty"`
	Credits int    `json:"credits,omitempty"`
}

type Course struct {
	database.Model
	Name            string            `gorm:"not null" json:"name"`
	Code            string            `gorm:"not null" json:"code"`
	Credits         int               `gorm:"not null" json:"credits"`
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
