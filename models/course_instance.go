package models

import (
	"uni_app/database"
)

type FetchCourseInstanceRequest struct {
	FetchRequest
	CourseID  database.PID `json:"course_id,omitempty"`
	TeacherID database.PID `json:"teacher_id,omitempty"`
	Semester  string       `json:"semester,omitempty"`
	Year      int          `json:"year,omitempty"`
}

type CourseInstance struct {
	database.Model
	CourseID       database.PID     `gorm:"not null" json:"course_id"`
	Course         *Course          `json:"course,omitempty"`
	TeacherID      database.PID     `gorm:"not null" json:"teacher_id"`
	Teacher        *Teacher         `json:"teacher,omitempty"`
	Students       []*Student       `gorm:"many2many:course_instance_students;" json:"students,omitempty"`
	ClassSchedules []*ClassSchedule `json:"class_schedules,omitempty"`
	Semester       string           `gorm:"not null" json:"semester"`
	Year           int              `gorm:"not null" json:"year"`
}
