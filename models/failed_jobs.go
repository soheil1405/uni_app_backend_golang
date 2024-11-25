package models

import (
	"time"

	"gorm.io/gorm"
)

type FailedJob struct {
	gorm.Model
	UUID       string    `gorm:"unique;not null"`
	Connection string    `gorm:"not null"`
	Queue      string    `gorm:"not null"`
	Payload    string    `gorm:"type:longtext;not null"`
	Exception  string    `gorm:"type:longtext;not null"`
	FailedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
