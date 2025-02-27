package models

import "uni_app/database"

type DaneshKadeha []*DaneshKadeha

type DaneshKadeh struct {
	database.Model
	Name              string           `json:"name,omitempty"`
	Description       string           `json:"description,omitempty"`
	UniID             database.PID     `json:"uni_id,omitempty"`
	Uni               *Uni             `gorm:"auto_preload:false;foreignKey:UniID;constraint:OnDelete:CASCADE;" json:"uni,omitempty"`
	Phones            []Phone          `gorm:"polymorphic:Owner;auto_preload:false;" json:"phones,omitempty"`
	Address           Address          `gorm:"polymorphic:Owner;auto_preload:false;" json:"address,omitempty"`
	ContactWays       []ContactWay     `gorm:"polymorphic:Owner;auto_preload:false;" json:"contact_ways,omitempty"`
	Students          []Student        `gorm:"foreignKey:DaneshKadehID;constraint:OnDelete:CASCADE;" json:"students,omitempty"`
	DaneshKadehTypeID database.PID     `json:"daneshkadeh_type_id,omitempty"`
	DaneshKadehType   *DaneshKadehType `gorm:"auto_preload:false;foreignKey:DaneshKadehTypeID;constraint:OnDelete:SET NULL;" json:"daneshkadeh_type,omitempty"`
	UserRoles         []*UserRole      `gorm:"auto_preload:false;foreignKey:DaneshKadehID;constraint:OnDelete:CASCADE;" json:"user_roles,omitempty"`
	MajorsCharts      MajorsCharts     `json:"majors_charts,omitempty" gorm:"foreignKey:DaneshKadehID;constraint:OnDelete:CASCADE;"`
}

type DaneshKadehType struct {
	database.Model
	Name        string `json:"name,omitempty"`
	Type        string `gorm:"unique;not null" json:"type,omitempty"`
	Description string `gorm:"type:text" json:"description,omitempty"`
}
