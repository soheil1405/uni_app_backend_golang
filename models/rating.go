package models

import (
	"time"
	"uni_app/database"
)

// Rating represents a polymorphic rating system that can be used to rate any entity
type Rating struct {
	ID        database.PID `json:"id" gorm:"primaryKey"`
	StudentID database.PID `json:"student_id" gorm:"not null"`
	Rating    float64      `json:"rating" gorm:"not null"`
	Comment   string       `json:"comment"`
	OwnerID   database.PID `json:"owner_id" gorm:"not null"`
	OwnerType string       `json:"owner_type" gorm:"not null"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	Student   *Student     `json:"student,omitempty" gorm:"foreignKey:StudentID"`
}

// FetchRatingRequest represents the request parameters for fetching ratings
type FetchRatingRequest struct {
	StudentID database.PID `json:"student_id" query:"student_id"`
	OwnerID   database.PID `json:"owner_id" query:"owner_id"`
	OwnerType string       `json:"owner_type" query:"owner_type"`
	MinRating float64      `json:"min_rating" query:"min_rating"`
	MaxRating float64      `json:"max_rating" query:"max_rating"`
	Includes  []string     `json:"includes" query:"includes"`
}

// Constants for different types of ratable entities
const (
	OwnerTypeUni         = "uni"
	OwnerTypeMajor       = "major"
	OwnerTypeLesson      = "lesson"
	OwnerTypeDaneshKadeh = "daneshkadeh"
	OwnerTypePlace       = "place"
)
