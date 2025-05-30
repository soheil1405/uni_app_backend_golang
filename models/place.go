package models

import "uni_app/database"

// جا ها
type Place struct {
	database.Model
	Name        string       `gorm:"not null" json:"name"`
	CityID      database.PID `gorm:"not null" json:"city_id"`
	City        *City        `gorm:"foreignKey:CityID" json:"city"`
	PlaceTypeID database.PID `gorm:"not null" json:"place_type_id"`
	PlaceType   *PlaceType   `gorm:"foreignKey:PlaceTypeID" json:"place_type"`
	AddressID   database.PID `gorm:"not null" json:"address_id"`
	Address     *Address     `gorm:"foreignKey:AddressID" json:"address"`
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
