package models

import "uni_app/database"

type UserRole struct {
	database.Model
	UserID        database.PID  `gorm:"not null" json:"user_id"`
	RoleID        database.PID  `gorm:"not null" json:"role_id"`
	DaneshKadehID *database.PID `json:"daneshkadeh_id,omitempty"`
	UniID         *database.PID `json:"uni_id,omitempty"`
	User          *User         `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"user,omitempty"`
	Role          *Role         `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE;" json:"role,omitempty"`
	DaneshKadeh   *DaneshKadeh  `gorm:"foreignKey:DaneshKadehID;constraint:OnDelete:SET NULL;" json:"daneshkadeh,omitempty"`
	Uni           *Uni          `gorm:"foreignKey:UniID;constraint:OnDelete:SET NULL;" json:"uni,omitempty"`
	Meta          string        `gorm:"type:json" json:"meta,omitempty"`
	Description   string        `json:"description,omitempty"`
}
