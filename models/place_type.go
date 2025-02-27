package models

import "uni_app/database"

// انواع جاها
type PlaceType struct {
	database.Model
	Type        string `gorm:"unique;not null"`
	Description string `gorm:"type:text"`
}
