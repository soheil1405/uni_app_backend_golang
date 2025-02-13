package models

import "uni_app/database"

type PlaceType struct {
	database.Model
	Type        string `gorm:"unique;not null"`
	Description string `gorm:"type:text"`
}
