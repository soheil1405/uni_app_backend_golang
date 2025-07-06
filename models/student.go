package models

import "uni_app/database"

type Students []*Student
type StudentStatus int

const (
	StudentStatusPending StudentStatus = iota
	StudentStatusActive
	StudentStatusInactive
)

type Student struct {
	database.Model
	Name            string                `json:"name,omitempty"`
	LastName        string                `json:"last_name,omitempty"`
	StudentCode     database.PID          `json:"student_code,omitempty"`
	NationalCode    database.PID          `json:"national_code,omitempty"`
	Password        string                `json:"-,omitempty"`
	Status          StudentStatus         `json:"status,omitempty" gorm:"default:1"`
	UniID           database.PID          `json:"uni_id,omitempty"`
	Uni             *Uni                  `gorm:"foreignKey:UniID" json:"uni,omitempty"`
	MajorID         database.PID          `json:"major_id,omitempty"`
	Major           *Major                `json:"major,omitempty" gorm:"foreignKey:MajorID"`
	DegreeLevel     DegreeLevel           `json:"degree_level,omitempty"`
	FacultyID       database.PID          `json:"faculty_id,omitempty"`
	Faculty         *Faculty              `json:"faculty,omitempty" gorm:"foreignKey:FacultyID"`
	StudentNumber   string                `gorm:"uniqueIndex" json:"student_number,omitempty"`
	EntryYear       int                   `json:"entry_year,omitempty"`
	EntryTerm       int                   `json:"entry_term,omitempty"`
	CurrentTerm     int                   `json:"current_term,omitempty"`
	PassedCourses   []StudentPassedCourse `json:"passed_courses,omitempty" gorm:"foreignKey:StudentID;constraint:OnDelete:CASCADE;"`
	CourseInstances []*CourseInstance     `json:"course_instances,omitempty" gorm:"many2many:course_instance_students;"`
}

type FetchStudentRequest struct {
	FetchRequest
	UniID         database.PID  `json:"uni_id" query:"uni_id"`
	MajorID       database.PID  `json:"major_id" query:"major_id"`
	FacultyID     database.PID  `json:"Faculty_id" query:"Faculty_id"`
	StudentNumber string        `json:"student_number" query:"student_number"`
	EntryYear     int           `json:"entry_year" query:"entry_year"`
	EntryTerm     int           `json:"entry_term" query:"entry_term"`
	CurrentTerm   int           `json:"current_term" query:"current_term"`
	Status        StudentStatus `json:"status" query:"status"`
}

func StudentAcceptIncludes() []string {
	return []string{
		"Uni",
		"Major",
		"Faculty",
		"PassedCourses",
		"CurrentCourses",
	}
}
