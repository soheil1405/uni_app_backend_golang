package models

import "uni_app/database"

type DaneshKadehType string

const (
	DaneshKadehTypeFanni    DaneshKadehType = "fanni"
	DaneshKadehTypeEnsani   DaneshKadehType = "ensani"
	DaneshKadehTypeMemari   DaneshKadehType = "memari"
	DaneshKadehTypeModiriat DaneshKadehType = "modiriat"
	DaneshKadehTypeHonar    DaneshKadehType = "honar"
)

type DaneshKadeha []*DaneshKadeh

type DaneshKadeh struct {
	database.Model
	Name            string           `json:"name,omitempty"`
	Description     string           `json:"description,omitempty"`
	UniID           database.PID     `json:"uni_id,omitempty"`
	Uni             *Uni             `gorm:"auto_preload:false;foreignKey:UniID;constraint:OnDelete:CASCADE;" json:"uni,omitempty"`
	Phones          []Phone          `gorm:"polymorphic:Owner;auto_preload:false;" json:"phones,omitempty"`
	Address         Address          `gorm:"polymorphic:Owner;auto_preload:false;" json:"address,omitempty"`
	ContactWays     []ContactWay     `gorm:"polymorphic:Owner;auto_preload:false;" json:"contact_ways,omitempty"`
	Students        []Student        `gorm:"foreignKey:DaneshKadehID;constraint:OnDelete:CASCADE;" json:"students,omitempty"`
	DaneshKadehType *DaneshKadehType `json:"daneshkadeh_type,omitempty"`
	UserRoles       []*UserRole      `gorm:"auto_preload:false;foreignKey:DaneshKadehID;constraint:OnDelete:CASCADE;" json:"user_roles,omitempty"`
	MajorsCharts    MajorsCharts     `json:"majors_charts,omitempty" gorm:"foreignKey:DaneshKadehID;constraint:OnDelete:CASCADE;"`
	Ratings         []Rating         `json:"ratings,omitempty" gorm:"polymorphic:Owner;polymorphicValue:daneshkadeh"`
}

type FetchDaneshKadehRequest struct {
	FetchRequest
	UniID           database.PID    `json:"uni_id" query:"uni_id"`
	DaneshKadehType DaneshKadehType `json:"daneshkadeh_type" query:"daneshkadeh_type"`
}

func DaneshKadehAcceptIncludes() []string {
	return []string{
		"Uni",
		"Phones",
		"Address",
		"ContactWays",
		"Students",
		"DaneshKadehType",
		"UserRoles",
		"MajorsCharts",
		"Ratings",
	}
}
