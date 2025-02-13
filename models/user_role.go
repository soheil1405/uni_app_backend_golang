package models

import "uni_app/database"

type UserRole struct {
	database.Model
	UserID database.PID
	User   User `gorm:"foreignKey:UserID"`
	RoleID database.PID
	Role   Role   `gorm:"foreignKey:RoleID"`
	Meta   string `gorm:"type:json"`
}
