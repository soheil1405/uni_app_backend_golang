package models

import "uni_app/database"

type Addresses []*Address

type Address struct {
	database.Model

	Title       string       `json:"title,omitempty"`
	FullAddress string       `json:"full_address,omitempty"`
	Street      string       `json:"street,omitempty"`
	Pelak       string       `json:"pelak,omitempty"`
	PostalCode  *int         `json:"postal_code,omitempty"`
	Phones      Phones       `json:"phones,omitempty"`
	CityID      database.PID `json:"city_id,omitempty"`
	City        City         `json:"city,omitempty"`
	Location    Location     `json:"location,omitempty"`

	PolymorphicModel
}

type Location struct {
	database.Model
	PolymorphicModel
	Lat float64 `json:"lat,omitempty"`
	Lng float64 `json:"lng,omitempty"`
}
