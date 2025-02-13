package models

import (
	"time"
	"uni_app/database"
)

type Uni struct {
	database.Model
	Name            string
	ZipCode         *int
	Website         *string
	Email           *string
	PhoneNumber1    *string
	PhoneNumber2    *string
	UniTypeID       uint
	UniType         UniType `gorm:"foreignKey:UniTypeID"`
	BossID          uint
	Boss            User `gorm:"foreignKey:BossID"`
	CityID          uint
	City            City `gorm:"foreignKey:CityID"`
	EstablishedYear *time.Time
}
