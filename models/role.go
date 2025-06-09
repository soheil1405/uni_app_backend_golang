package models

import (
	"encoding/json"
	"uni_app/database"

	"gorm.io/gorm"
)

type Roles []Role

// Role model

type Role struct {
	database.Model
	Name        string          `gorm:"not null" json:"name"`
	Priority    int             `gorm:"not null;unique;" json:"priority"`
	Description string          `json:"description,omitempty"`
	Meta        json.RawMessage `gorm:"type:json" json:"meta,omitempty"`
	UserRoles   []*UserRole     `json:"user_roles,omitempty"`
}

func (role *Role) BeforeCreate(tx *gorm.DB) (err error) {
	var maxPriority *int
	if err := tx.Model(&Role{}).Select("MAX(priority)").Scan(&maxPriority).Error; err != nil {
		return err
	}

	if role.Priority == 0 {
		if maxPriority == nil {
			role.Priority = 1
		} else {
			role.Priority = *maxPriority + 1
		}
	}
	return
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
