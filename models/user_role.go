package models

import "uni_app/database"

type UserRole struct {
	database.Model

	UserID database.PID `json:"user_id,omitempty"`
	User   User         `gorm:"foreignKey:UserID" json:"user,omitempty"`

	RoleID database.PID `json:"role_id,omitempty"`
	Role   Role         `gorm:"foreignKey:RoleID" json:"role,omitempty"`

	UniID database.PID `json:"uni_id,omitempty"`
	Uni   Uni          `gorm:"foreignKey:UniID" json:"uni,omitempty"`

	DaneshKadehID database.PID `json:"daneshkadeh_id,omitempty"`
	DaneshKadeh   DaneshKadeh  `gorm:"foreignKey:daneshkadehID" json:"daneshkadeh,omitempty"`

	Meta string `gorm:"type:json" json:"meta,omitempty"`
}
