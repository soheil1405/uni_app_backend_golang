package models

import "uni_app/database"

type ClassTimes []*ClassTime

type ClassTime struct {
	database.Model
	StudentCurrentLessonID database.PID          `json:"student_current_lesson_id,omitempty"`
	StudentCurrentLesson   *StudentCurrentLesson `gorm:"foreignKey:StudentCurrentLessonID;constraint:OnDelete:CASCADE;" json:"student_current_lesson,omitempty"`
	DayOfWeek              int                   `json:"day_of_week,omitempty"` // 0-6 (Sunday-Saturday)
	StartTime              string                `json:"start_time,omitempty"`  // Format: "HH:MM"
	EndTime                string                `json:"end_time,omitempty"`    // Format: "HH:MM"
	ClassroomNumber        string                `json:"classroom_number,omitempty"`
	Building               string                `json:"building,omitempty"`
	IsSelected             bool                  `gorm:"default:false" json:"is_selected,omitempty"`
	Description            string                `json:"description,omitempty"`
}

type FetchClassTimeRequest struct {
	FetchRequest
	StudentCurrentLessonID database.PID `json:"student_current_lesson_id" query:"student_current_lesson_id"`
	DayOfWeek              int          `json:"day_of_week" query:"day_of_week"`
	IsSelected             bool         `json:"is_selected" query:"is_selected"`
}

func ClassTimeAcceptIncludes() []string {
	return []string{
		"StudentCurrentLesson",
	}
}
