package models

import "uni_app/database"

type Addresses []*Address

type Address struct {
	database.Model
	PolymorphicModel
	Title       string       `json:"title,omitempty"`
	FullAddress string       `json:"full_address,omitempty"`
	Street      string       `json:"street,omitempty"`
	Pelak       string       `json:"pelak,omitempty"`
	PostalCode  *int         `json:"postal_code,omitempty"`
	Phones      Phones       `json:"phones,omitempty" gorm:"polymorphic:Owner;"`
	CityID      database.PID `json:"city_id,omitempty"`
	City        City         `json:"city,omitempty" gorm:"foreignKey:CityID"`
	Lat         float64      `json:"lat,omitempty"`
	Lng         float64      `json:"lng,omitempty"`
}

type FetchAddressRequest struct {
	FetchRequest
	CityID database.PID `json:"city_id" query:"city_id"`
}

func AddressAcceptIncludes() []string {
	return []string{
		"Phones",
		"City",
	}
}
