package models

import (
	"time"
	"uni_app/database"
)

type FetchCourseInstanceRequest struct {
	FetchRequest
	CourseID  database.PID `json:"course_id,omitempty"`
	TeacherID database.PID `json:"teacher_id,omitempty"`
}

type CourseInstance struct {
	database.Model
	UniID          database.PID `json:"uni_id,omitempty"`
	Uni            Uni          `json:"uni,omitempty"`
	CourseID       database.PID `gorm:"not null" json:"course_id,omitempty"`
	Course         *Course      `json:"course,omitempty"`
	TeacherID      database.PID `gorm:"not null" json:"teacher_id,omitempty"`
	Teacher        *Teacher     `json:"teacher,omitempty"`
	Students       []*Student   `gorm:"many2many:course_instance_students;" json:"students,omitempty"`
	FacultyID      database.PID `json:"faculty_id,omitempty"`
	Faculty        Faculty      `json:"faculty,omitempty"`
	RoomID         database.PID `gorm:"not null" json:"room_id,omitempty"`
	Room           *Room        `json:"room,omitempty"`
	DayOfWeek      int          `gorm:"not null" json:"day_of_week,omitempty"`
	StartTime      time.Time    `gorm:"not null" json:"start_time,omitempty"`
	EndTime        time.Time    `gorm:"not null" json:"end_time,omitempty"`
	FirstDayInTerm time.Time    `json:"first_day_in_term,omitempty"`
	LastDayInTerm  time.Time    `json:"last_day_in_term,omitempty"`
	ExamTime       time.Time    `json:"exam_time,omitempty"`
	FinishedAt     time.Time    `json:"finished_at,omitempty"`
}
