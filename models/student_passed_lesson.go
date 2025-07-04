package models

import (
	"time"
	"uni_app/database"
)

// StudentPassedCourse represents a passed Course for a student
type StudentPassedCourse struct {
	database.Model
	StudentID     database.PID `json:"student_id" gorm:"not null"`
	Student       Student      `json:"student" gorm:"foreignKey:StudentID"`
	CourseID      database.PID `json:"course_id,omitempty"`
	Course        *Course      `gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE;" json:"course,omitempty"`
	Grade         float64      `json:"grade" gorm:"type:decimal(4,2)"`
	Term          int          `json:"term,omitempty"`
	TermStartTime time.Time    `json:"term_start_time,omitempty"`
	TermEndTime   time.Time    `json:"term_end_time,omitempty"`
	ExamTime      time.Time    `json:"exam_time,omitempty"`
}

func StudentPassedCourseAcceptIncludes() []string {
	return []string{
		"Student",
		"Course",
	}
}

type FetchStudentPassedCourseRequest struct {
	FetchRequest
	StudentID database.PID `json:"student_id" query:"student_id"`
	CourseID  database.PID `json:"course_id" query:"course_id"`
	Term      int          `json:"term" query:"term"`
}
