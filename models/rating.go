package models

import (
	"uni_app/database"
)

// Rating represents a polymorphic rating system that can be used to rate any entity
type Rating struct {
	database.Model
	PolymorphicModel
	StudentID database.PID `json:"student_id" gorm:"not null"`
	Student   *Student     `json:"student,omitempty" gorm:"foreignKey:StudentID"`
	Rate      int          `json:"rate" gorm:"not null"`
}

// FetchRatingRequest represents the request parameters for fetching ratings
type FetchRatingRequest struct {
	StudentID database.PID  `json:"student_id" query:"student_id"`
	UserID    database.PID  `json:"user_id" query:"user_id"`
	OwnerID   database.PID  `json:"owner_id" query:"owner_id"`
	OwnerType string        `json:"owner_type" query:"owner_type"`
	ParentID  *database.PID `json:"parent_id" query:"parent_id"`
	MinRating float64       `json:"min_rating" query:"min_rating"`
	MaxRating float64       `json:"max_rating" query:"max_rating"`
	Includes  []string      `json:"includes" query:"includes"`
}

// Constants for different types of ratable entities
const (
	OwnerTypeUni         = "uni"
	OwnerTypeMajor       = "major"
	OwnerTypeLesson      = "lesson"
	OwnerTypeDaneshKadeh = "daneshkadeh"
	OwnerTypePlace       = "place"
)
