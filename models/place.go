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

type FetchPlaceRequest struct {
	FetchRequest
	CityID      database.PID `json:"city_id" query:"city_id"`
	PlaceTypeID database.PID `json:"place_type_id" query:"place_type_id"`
}

func PlaceAcceptIncludes() []string {
	return []string{
		"City",
		"PlaceType",
	}
}
