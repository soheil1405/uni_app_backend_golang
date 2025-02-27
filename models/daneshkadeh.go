package models

import "uni_app/database"

type DaneshKadeha []*DaneshKadeha

type DaneshKadeh struct {
	database.Model
	Name              string           `json:"name"`
	Description       string           `json:"description,omitempty"`
	UniID             database.PID     `json:"uni_id,omitempty"`
	Uni               *Uni             `gorm:"foreignKey:UniID;constraint:OnDelete:CASCADE;" json:"uni,omitempty"`
	Phones            []Phone          `gorm:"polymorphic:Owner;" json:"phones,omitempty"`
	Address           Address          `gorm:"polymorphic:Owner;" json:"address,omitempty"`
	ContactWays       []ContactWay     `gorm:"polymorphic:Owner;" json:"contact_ways,omitempty"`
	DaneshKadehTypeID database.PID     `json:"daneshkadeh_type_id,omitempty"`
	DaneshKadehType   *DaneshKadehType `gorm:"foreignKey:DaneshKadehTypeID;constraint:OnDelete:SET NULL;" json:"daneshkadeh_type,omitempty"`
	UserRoles         []*UserRole      `gorm:"foreignKey:DaneshKadehID;constraint:OnDelete:CASCADE;" json:"user_roles,omitempty"`
}

type DaneshKadehType struct {
	database.Model
	Type        string `gorm:"unique;not null"`
	Description string `gorm:"type:text"`
}
