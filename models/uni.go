package models

import (
	"time"
	"uni_app/database"
)

type Unis []*Uni
type Uni struct {
	database.Model
	Name         string       `json:"name,omitempty"`
	DaneshKadeha DaneshKadeha `json:"danesh_kadeha,omitempty"`
	UniTypeID    database.PID `json:"uni_type_id,omitempty"`
	UniType      UniType      `json:"uni_type,omitempty" gorm:"foreignKey:UniTypeID"`
	Students     Students     `json:"students,omitempty"`
	Users        Users        `gorm:"many2many:user_roles;" json:"users,omitempty"`
	Roles        Roles        `gorm:"many2many:user_roles;" json:"roles,omitempty"`
	Phones       Phones       `json:"phones,omitempty"`
	Address      Address      `json:"address,omitempty"`
	ContactWays  ContactWays  `json:"contact_ways,omitempty"`

	// UserRoles       []*UserRole  `json:"user_roles,omitempty" gorm:"association_autoupdate:false;association_autocreate:false;"`
	EstablishedYear *time.Time `json:"established_year,omitempty"`
}
