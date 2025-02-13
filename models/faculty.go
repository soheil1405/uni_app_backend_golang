package models

import "uni_app/database"

type Faculty struct {
	database.Model
	UniversityID database.PID
	University   Uni `gorm:"foreignKey:UniversityID"`
	Name         string
	Description  string
	BossID       uint
	Boss         User `gorm:"foreignKey:BossID"`
}
