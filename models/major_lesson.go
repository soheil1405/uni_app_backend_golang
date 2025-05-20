package models

import (
	"uni_app/database"
)

// MajorLesson represents a relationship between a major and a lesson
type MajorLesson struct {
	database.Model
	MajorChartID    uint        `gorm:"not null" json:"major_chart_id"`
	MajorChart      MajorsChart `gorm:"foreignKey:MajorChartID" json:"major_chart"`
	MajorID         uint        `gorm:"not null" json:"major_id"`
	Major           Major       `gorm:"foreignKey:MajorID" json:"major"`
	LessonID        uint        `gorm:"not null" json:"lesson_id"`
	Lesson          Lesson      `gorm:"foreignKey:LessonID" json:"lesson"`
	RecommendedTerm int         `gorm:"not null" json:"recommended_term"`
	IsOptional      bool        `gorm:"not null" json:"is_optional"`
	IsTechnical     bool        `gorm:"not null" json:"is_technical"`
	Ratio           float64     `gorm:"not null" json:"ratio"`
}
