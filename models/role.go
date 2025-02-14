package models

import "uni_app/database"
//نقش
type Role struct {
	database.Model
	Meta        string `gorm:"type:json"`
	Name        string `gorm:"unique;not null"`
	Description string
}
