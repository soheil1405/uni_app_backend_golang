package models

import (
	"time"
	"uni_app/database"
)

type UniType string

const (
	UniTypePezeshki UniType = "pezeshki"
)

type Unis []*Uni
type Uni struct {
	database.Model
	Name            string        `json:"name"`
	UniType         *UniType      `json:"uni_type,omitempty"`
	EstablishedYear *time.Time    `json:"established_year,omitempty"`
	CityID          database.PID  `json:"city_id,omitempty"`
	City            *City         `gorm:"foreignKey:CityID;constraint:OnDelete:SET NULL;" json:"city,omitempty"`
	Address         Address       `gorm:"polymorphic:Owner;" json:"addresses,omitempty"`
	Phones          []*Phone      `gorm:"polymorphic:Owner;" json:"phones,omitempty"`
	ContactWays     []*ContactWay `gorm:"polymorphic:Owner;" json:"contact_ways,omitempty"`
	Students        []*Student    `gorm:"foreignKey:UniID;constraint:OnDelete:CASCADE;" json:"students,omitempty"`
	Faculties       []*Faculty    `gorm:"foreignKey:UniID;constraint:OnDelete:CASCADE;" json:"daneshkadeha,omitempty"`
	Ratings         []*Rating     `json:"ratings,omitempty" gorm:"polymorphic:Owner;polymorphicValue:unis"`
	Teachers        []*Teacher    `json:"teachers,omitempty" gorm:"many2many:uni_teachers;"`
	Users           Users         `json:"users,omitempty" gorm:"foreignKey:UniID;constraint:OnDelete:CASCADE;"`
}

type FetchUniRequest struct {
	FetchRequest
	UniTypeID database.PID `json:"uni_type_id" query:"uni_type_id"`
	CityID    database.PID `json:"city_id" query:"city_id"`
}

func UniAcceptIncludes() []string {
	return []string{
		"UniType",
		"ContactWays",
		"Phones",
		"City",
		"Addresses",
		"Students",
		"DaneshKadeha",
		"UserRoles",
		"Ratings",
	}
}
