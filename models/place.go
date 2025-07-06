package models

import (
	"uni_app/database"
)

// جا ها
type Place struct {
	database.Model
	OpenTime    int          `json:"open_time,omitempty"`
	CloseTime   int          `json:"close_time,omitempty"`
	Active      bool         `json:"active"`
	Name        string       `gorm:"not null" json:"name,omitempty"`
	CityID      database.PID `gorm:"not null" json:"city_id,omitempty"`
	City        *City        `gorm:"foreignKey:CityID" json:"city,omitempty"`
	PlaceTypeID database.PID `gorm:"not null" json:"place_type_id,omitempty"`
	PlaceType   *PlaceType   `gorm:"foreignKey:PlaceTypeID" json:"place_type,omitempty"`
	AddressID   database.PID `gorm:"not null" json:"address_id,omitempty"`
	Address     *Address     `gorm:"foreignKey:AddressID" json:"address,omitempty"`
	Ratings     []Rating     `json:"ratings,omitempty" gorm:"polymorphic:Owner;polymorphicValue:place"`
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
		"Ratings",
	}
}
