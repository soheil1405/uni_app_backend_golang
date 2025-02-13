package models

import "uni_app/database"

type UniType struct {
	database.Model
	Type        string `gorm:"unique;not null"`
	Description string `gorm:"type:text"`
}
