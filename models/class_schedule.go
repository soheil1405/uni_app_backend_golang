package models

import (
	"uni_app/database"
)

type FetchClassScheduleRequest struct {
	FetchRequest
	CourseInstanceID database.PID `json:"course_instance_id,omitempty"`
	RoomID           database.PID `json:"room_id,omitempty"`
	DayOfWeek        int          `json:"day_of_week,omitempty"`
	StartTime        string       `json:"start_time,omitempty"`
	EndTime          string       `json:"end_time,omitempty"`
}

type ClassSchedule struct {
	database.Model
	CourseInstanceID database.PID    `gorm:"not null" json:"course_instance_id"`
	CourseInstance   *CourseInstance `json:"course_instance,omitempty"`
	RoomID           database.PID    `gorm:"not null" json:"room_id"`
	Room             *Room           `json:"room,omitempty"`
	DayOfWeek        int             `gorm:"not null" json:"day_of_week"`
	StartTime        string          `gorm:"not null" json:"start_time"`
	EndTime          string          `gorm:"not null" json:"end_time"`
}
