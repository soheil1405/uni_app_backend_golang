package models

import (
	"uni_app/database"
)

type Roles []Role

// Role model

type Role struct {
	database.Model
	Name        string      `gorm:"not null" json:"name"`
	Priority    int         `gorm:"not null;unique;" json:"priority"`
	Description string      `json:"description,omitempty"`
	Meta        string      `gorm:"type:json" json:"meta,omitempty"`
	UserRoles   []*UserRole `json:"user_roles,omitempty"`
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

func RoleAcceptIncludes() []string {
	return []string{
		"UserRoles",
	}
}
