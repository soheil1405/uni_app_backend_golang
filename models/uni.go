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
	Active          bool          `json:"active"`
	Name            string        `json:"name"`
	UniType         *UniType      `json:"uni_type,omitempty"`
	EstablishedYear *time.Time    `json:"established_year,omitempty"`
	CityID          database.PID  `json:"city_id,omitempty"`
	City            *City         `json:"city,omitempty" gorm:"foreignKey:CityID;constraint:OnDelete:SET NULL;"`
	Address         Address       `json:"addresses,omitempty" gorm:"polymorphic:Owner;"`
	Dimains         []*Domain     `json:"domains,omitempty"  gorm:"foreignKey:UniID;constraint:OnDelete:CASCADE;"`
	Phones          []*Phone      `json:"phones,omitempty"  gorm:"polymorphic:Owner;"`
	ContactWays     []*ContactWay `json:"contact_ways,omitempty"  gorm:"polymorphic:Owner;"`
	Students        []*Student    `json:"students,omitempty"  gorm:"foreignKey:UniID;constraint:OnDelete:CASCADE;"`
	Faculties       []*Faculty    `json:"daneshkadeha,omitempty"  gorm:"foreignKey:UniID;constraint:OnDelete:CASCADE;"`
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
		"Dimains",
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
