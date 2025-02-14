package models

import "uni_app/database"

type Faculties []*Faculty

// دانشکده
type Faculty struct {
	database.Model
	UniversityID database.PID
	University   Uni `gorm:"foreignKey:UniversityID"`
	Name         string
	Description  string
	BossID       database.PID
	Boss         User `gorm:"foreignKey:BossID"`
}
