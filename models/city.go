package models

import "uni_app/database"
//شهر
type City struct {
	database.Model
	Name  string `gorm:"not null"`
	Unies Unies
}
