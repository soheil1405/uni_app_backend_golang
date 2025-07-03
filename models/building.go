package models

import (
	"uni_app/database"
)

type FetchBuildingRequest struct {
	FetchRequest
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}

type Building struct {
	database.Model
	Name    string  `gorm:"not null" json:"name"`
	Address string  `gorm:"not null" json:"address"`
	Rooms   []*Room `json:"rooms,omitempty"`
}
