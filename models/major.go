package models

import "uni_app/database"

type Major struct {
	database.Model
	Name        string `gorm:"not null"`
	Code        string `gorm:"unique;not null"`
	Description string
}
