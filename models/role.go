package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Meta        string `gorm:"type:json"`
	Name        string `gorm:"unique;not null"`
	Description string
}
