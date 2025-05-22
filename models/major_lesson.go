package models

import (
	"uni_app/database"
)

// MajorLesson represents a relationship between a major and a lesson
type MajorLesson struct {
	database.Model
	MajorChartID    database.PID `gorm:"not null" json:"major_chart_id"`
	MajorChart      MajorsChart  `gorm:"foreignKey:MajorChartID" json:"major_chart"`
	MajorID         database.PID `gorm:"not null" json:"major_id"`
	Major           Major        `gorm:"foreignKey:MajorID" json:"major"`
	LessonID        database.PID `gorm:"not null" json:"lesson_id"`
	Lesson          Lesson       `gorm:"foreignKey:LessonID" json:"lesson"`
	RecommendedTerm int          `gorm:"not null" json:"recommended_term"`
	IsOptional      bool         `gorm:"not null" json:"is_optional"`
	IsTechnical     bool         `gorm:"not null" json:"is_technical"`
	Ratio           float64      `gorm:"not null" json:"ratio"`
}

type FetchMajorLessonRequest struct {
	FetchRequest
	MajorID         database.PID `json:"major_id" query:"major_id"`
	MajorChartID    database.PID `json:"major_chart_id" query:"major_chart_id"`
	LessonID        database.PID `json:"lesson_id" query:"lesson_id"`
	IsOptional      bool         `json:"is_optional" query:"is_optional"`
	IsTechnical     bool         `json:"is_technical" query:"is_technical"`
	RecommendedTerm int          `json:"recommended_term" query:"recommended_term"`
}

func MajorLessonAcceptIncludes() []string {
	return []string{
		"MajorChart",
		"Major",
		"Lesson",
	}
}
