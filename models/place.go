package models

import "uni_app/database"

type Place struct {
	database.Model
	Name        string       `gorm:"not null"`
	CityID      uint         `gorm:"foreignKey:ID"`
	City        City         `gorm:"foreignKey:CityID"`
	PlaceTypeID database.PID `gorm:"foreignKey:ID"`
	PlaceType   PlaceType    `gorm:"foreignKey:PlaceTypeID"`
}
