package models

import (
	"uni_app/database"
)

// UserMode ...
type UserMode uint

const (
	// PENDING ...
	PENDING UserMode = iota + 1
	// USER_STATUS_ACTIVE Active
	USER_STATUS_ACTIVE
	// USER_STATUS_INACTIVE Inactive
	USER_STATUS_INACTIVE
)

type Users []*User

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
	Status        UserMode     `json:"status" gorm:"default:1"`

	Roles        Roles        `gorm:"many2many:user_roles;" json:"roles,omitempty"`
	Unis         Unis         `json:"unis,omitempty" gorm:"many2many:user_roles;"`
	DaneshKadeha DaneshKadeha `json:"daneshKadeha,omitempty" gorm:"many2many:user_roles;"`

	// UserRoles     []*UserRole  `json:"user_roles,omitempty" gorm:"association_autoupdate:false;association_autocreate:false;"`
}

type UserLoginRequst struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
