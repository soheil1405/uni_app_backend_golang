package models

import (
	"uni_app/database"
)

// StudentPassedLesson represents a passed lesson for a student
type StudentPassedLesson struct {
	database.Model
	StudentID database.PID `json:"student_id" gorm:"not null"`
	Student   Student      `json:"student" gorm:"foreignKey:StudentID"`
	LessonID  database.PID `json:"lesson_id" gorm:"not null"`
	Lesson    Lesson       `json:"lesson" gorm:"foreignKey:LessonID"`
	Grade     float64      `json:"grade" gorm:"type:decimal(4,2)"`
	Term      int          `json:"term" gorm:"not null"`
}

func StudentPassedLessonAcceptIncludes() []string {
	return []string{
		"Student",
		"Lesson",
	}
}

type FetchStudentPassedLessonRequest struct {
	FetchRequest
	StudentID database.PID `json:"student_id" query:"student_id"`
	LessonID  database.PID `json:"lesson_id" query:"lesson_id"`
	Term      int          `json:"term" query:"term"`
}
