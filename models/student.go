package models

import "uni_app/database"

type Students []*Student
type StudentStatus int

const (
	StudentStatusActive StudentStatus = iota
	StudentStatusInactive
	StudentStatusPending
)

type Student struct {
	database.Model
	Name           string                 `json:"name,omitempty"`
	LastName       string                 `json:"last_name,omitempty"`
	StudentCode    database.PID           `json:"student_code,omitempty"`
	NationalCode   database.PID           `json:"national_code,omitempty"`
	Password       string                 `json:"-"`
	Status         StudentStatus          `json:"status,omitempty" gorm:"default:1"`
	UniID          database.PID           `json:"uni_id,omitempty"`
	Uni            *Uni                   `gorm:"foreignKey:UniID" json:"uni,omitempty"`
	MajorID        database.PID           `json:"major_id,omitempty"`
	Major          *Major                 `json:"major,omitempty" gorm:"foreignKey:MajorID"`
	DaneshKadehID  database.PID           `json:"daneshkadeh_id,omitempty"`
	DaneshKadeh    *DaneshKadeh           `json:"danesh_kadeh,omitempty" gorm:"foreignKey:DaneshKadehID"`
	StudentNumber  string                 `gorm:"uniqueIndex" json:"student_number,omitempty"`
	EntryYear      int                    `json:"entry_year,omitempty"`
	EntryTerm      int                    `json:"entry_term,omitempty"`
	CurrentTerm    int                    `json:"current_term,omitempty"`
	PassedLessons  []StudentPassedLesson  `json:"passed_lessons,omitempty" gorm:"foreignKey:StudentID;constraint:OnDelete:CASCADE;"`
	CurrentLessons []StudentCurrentLesson `json:"current_lessons,omitempty" gorm:"foreignKey:StudentID;constraint:OnDelete:CASCADE;"`
}

type FetchStudentRequest struct {
	FetchRequest
	UniID         database.PID  `json:"uni_id" query:"uni_id"`
	MajorID       database.PID  `json:"major_id" query:"major_id"`
	DaneshKadehID database.PID  `json:"daneshkadeh_id" query:"daneshkadeh_id"`
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
		"DaneshKadeh",
		"PassedLessons",
		"CurrentLessons",
	}
}
