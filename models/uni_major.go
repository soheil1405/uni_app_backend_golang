package models

import (
	"gorm.io/gorm"
)

type UniMajor struct {
	gorm.Model
	DegreeLevel  string
	Duration     int
	Description  string
	UniversityID uint
	University   Uni `gorm:"foreignKey:UniversityID"`
	FacultyID    uint
	Faculty      Faculty `gorm:"foreignKey:FacultyID"`
	MajorID      uint
	Major        Major `gorm:"foreignKey:MajorID"`
}
