package models

import (
	"time"

	"gorm.io/gorm"
)

type PersonalAccessToken struct {
	gorm.Model
	TokenableID   uint
	TokenableType string
	Name          string `gorm:"not null"`
	Token         string `gorm:"unique;size:64;not null"`
	Abilities     string `gorm:"type:text"`
	LastUsedAt    *time.Time
	ExpiresAt     *time.Time
}
