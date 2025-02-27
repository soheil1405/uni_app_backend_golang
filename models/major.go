package models

import "uni_app/database"

// رشته تحصیلی
type Major struct {
	database.Model
	Name        string    `gorm:"not null" json:"name,omitempty"`
	Code        string    `gorm:"unique;not null" json:"code,omitempty"`
	Description string    `json:"description,omitempty"`
	Students    []Student `gorm:"foreignKey:MajorID;constraint:OnDelete:CASCADE;" json:"students,omitempty"`
}
