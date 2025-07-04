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
	FetchRequest `json:"fetch_request,omitempty"`
	Name         string      `json:"name,omitempty"`
	Email        string      `json:"email,omitempty"`
	Phone        string      `json:"phone,omitempty"`
	DegreeLevel  DegreeLevel `json:"degree_level,omitempty"`
}

type Teacher struct {
	database.Model
	Name            string            `gorm:"not null" json:"name"`
	Email           string            `gorm:"not null" json:"email"`
	Phone           string            `gorm:"not null" json:"phone"`
	Unis            Unis              `json:"teachers,omitempty" gorm:"many2many:uni_teachers;"`
	DegreeLevel     DegreeLevel       `json:"degree_level,omitempty"`
	CourseInstances []*CourseInstance `json:"course_instances,omitempty" gorm:"foreignKey:TeacherID;constraint:OnDelete:CASCADE;"`
}
