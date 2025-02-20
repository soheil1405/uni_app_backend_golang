package models

import "uni_app/database"

type UniMajor struct {
	database.Model
	DegreeLevel  string
	Duration     int
	Description  string
	UniversityID database.PID
	University   Uni `gorm:"foreignKey:UniversityID"`
	FacultyID    database.PID
	Faculty      DaneshKadeha `gorm:"foreignKey:FacultyID"`
	MajorID      database.PID
	Major        Major `gorm:"foreignKey:MajorID"`
}
