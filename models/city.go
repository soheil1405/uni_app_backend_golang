package models

import "uni_app/database"

type City struct {
	database.Model
	Name string `gorm:"not null"`
}
