package models

import "uni_app/database"
//چارت رشته تحصیلی
type MajorsChart struct {
	database.Model
	Name       string `gorm:"not null"`
	UniMajorID database.PID
	UniMajor   UniMajor `gorm:"foreignKey:UniMajorID"`
}
