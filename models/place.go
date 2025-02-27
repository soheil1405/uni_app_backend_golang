package models

import "uni_app/database"

// جا ها
type Place struct {
	database.Model
	Name        string       `gorm:"not null"`
	CityID      database.PID `gorm:"not null"`
	City        *City        `gorm:"foreignKey:CityID"`
	PlaceTypeID database.PID `gorm:"not null"`
	PlaceType   *PlaceType   `gorm:"foreignKey:PlaceTypeID"`
}
