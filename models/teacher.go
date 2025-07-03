package models

import (
	"uni_app/database"
)

type TeacherTitle string

const (
	TeacherTitleProfessor     TeacherTitle = "professor"
	TeacherTitleAssociateProf TeacherTitle = "associate_professor"
	TeacherTitleAssistantProf TeacherTitle = "assistant_professor"
	TeacherTitleInstructor    TeacherTitle = "instructor"
)

type FetchTeacherRequest struct {
	FetchRequest
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type Teacher struct {
	database.Model
	Name            string            `gorm:"not null" json:"name"`
	Email           string            `gorm:"not null" json:"email"`
	Phone           string            `gorm:"not null" json:"phone"`
	CourseInstances []*CourseInstance `json:"course_instances,omitempty"`
}
