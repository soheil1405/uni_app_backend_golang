package models

import (
	"uni_app/database"
)

type Roles []Role

// نقش
type Role struct {
	database.Model
	Name        string `gorm:"not null;" json:"name,omitempty"`
	Priority    int    `gorm:"not null;defualt:1" json:"priority,omitempty"`
	Description string `json:"description,omitempty"`
	Meta        string `gorm:"type:json" json:"meta,omitempty"`
	// Users       Users       `gorm:"many2many:user_roles;" json:"users,omitempty"`
	UserRoles []*UserRole `json:"user_roles,omitempty" gorm:"association_autoupdate:false;association_autocreate:false;"`
}

func (roles Roles) GetMainRole() (role *Role) {
	if len(roles) == 0 {
		return nil
	}

	priority := 0
	for _, r := range roles {
		if r.Priority > priority {
			priority = r.Priority
			role = &r
		}
	}
	return
}
