package models

import "uni_app/database"
//رشته تحصیلی
type Major struct {
	database.Model
	Name        string `gorm:"not null"`
	Code        string `gorm:"unique;not null"`
	Description string
}
