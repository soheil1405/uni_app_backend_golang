package models

import "uni_app/database"

// نقش
type Role struct {
	database.Model
	Meta        string `gorm:"type:json" json:"meta,omitempty"`
	Name        string `gorm:"unique;not null" json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
