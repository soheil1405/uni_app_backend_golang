package models

import (
	"time"
	"uni_app/database"
)

type Unis []*Uni
type Uni struct {
	database.Model
	Name            string        `json:"name"`
	UniTypeID       database.PID  `json:"uni_type_id,omitempty"`
	UniType         *UniType      `gorm:"foreignKey:UniTypeID;constraint:OnDelete:SET NULL;" json:"uni_type,omitempty"`
	ContactWays     []ContactWay  `gorm:"polymorphic:Owner;" json:"contact_ways,omitempty"`
	EstablishedYear *time.Time    `json:"established_year,omitempty"`
	Phones          []Phone       `gorm:"polymorphic:Owner;" json:"phones,omitempty"`
	CityID          database.PID  `json:"city_id,omitempty"`
	City            *City         `gorm:"foreignKey:CityID;constraint:OnDelete:SET NULL;" json:"city,omitempty"`
	Addresses       []Address     `gorm:"polymorphic:Owner;" json:"addresses,omitempty"`
	Students        []Student     `gorm:"foreignKey:UniID;constraint:OnDelete:CASCADE;" json:"students,omitempty"`
	DaneshKadeha    []DaneshKadeh `gorm:"foreignKey:UniID;constraint:OnDelete:CASCADE;" json:"daneshkadeha,omitempty"`
	UserRoles       []*UserRole   `json:"user_roles,omitempty"`
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
	}
}
