package models

import (
	"uni_app/database"
)

type User struct {
	database.Model
	UserName      string       `gorm:"unique;not null" json:"user_name,omitempty"`
	FirstName     string       `gorm:"not null" json:"first_name,omitempty"`
	LastName      string       `gorm:"not null" json:"last_name,omitempty"`
	Number        string       `gorm:"unique;not null" json:"number,omitempty"`
	NationalCode  *int         `gorm:"unique" json:"national_code,omitempty"`
	NominatedByID database.PID `gorm:"foreignKey:ID" json:"nominated_by_id,omitempty"`
	NominatedBy   *User        `gorm:"foreignKey:NominatedByID" json:"nominated_by,omitempty"`
	Email         string       `gorm:"unique;not null" json:"email,omitempty"`
	Password      string       `gorm:"not null" json:"password,omitempty"`
	Roles         []Role       `gorm:"many2many:user_roles;" json:"roles,omitempty"`
}

type UserLoginRequst struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
