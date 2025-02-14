package models

import (
	"time"
	"uni_app/database"
)

type Unies []*Uni
type Uni struct {
	database.Model
	Name            string
	ZipCode         *int
	Website         *string
	Email           *string
	PhoneNumber1    *string
	PhoneNumber2    *string
	UniTypeID       database.PID
	UniType         UniType `gorm:"foreignKey:UniTypeID"`
	BossID          database.PID
	Boss            User `gorm:"foreignKey:BossID"`
	CityID          database.PID
	City            City `gorm:"foreignKey:CityID"`
	EstablishedYear *time.Time
	Faculties       Faculties
}
