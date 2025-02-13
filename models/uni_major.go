package models

import "uni_app/database"

type UniMajor struct {
	database.Model
	DegreeLevel  string
	Duration     int
	Description  string
	UniversityID database.PID
	University   Uni `gorm:"foreignKey:UniversityID"`
	FacultyID    uint
	Faculty      Faculty `gorm:"foreignKey:FacultyID"`
	MajorID      uint
	Major        Major `gorm:"foreignKey:MajorID"`
}
