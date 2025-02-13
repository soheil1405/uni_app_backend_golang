package models

import (
	"time"
	"uni_app/database"
)

type User struct {
	database.Model
	UserName        string `gorm:"unique;not null"`
	FirstName       string `gorm:"not null"`
	LastName        string `gorm:"not null"`
	Number          string `gorm:"unique;not null"`
	NationalCode    *int   `gorm:"unique"`
	NominatedByID   uint   `gorm:"foreignKey:ID"`
	NominatedBy     *User  `gorm:"foreignKey:NominatedByID"`
	Email           string `gorm:"unique;not null"`
	EmailVerifiedAt *time.Time
	Password        string `gorm:"not null"`
	RememberToken   string
}
