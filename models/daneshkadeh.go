package models

import "uni_app/database"

type DaneshKadeha []*DaneshKadeha

// دانشکده
type DaneshKadeh struct {
	database.Model
	Name         string
	Description  string
	University   Uni `gorm:"foreignKey:UniversityID"`
	UniversityID database.PID
	Boss         User `gorm:"foreignKey:BossID"`
	BossID       database.PID
}
