package models

import "uni_app/database"

type StudentRole struct {
	database.Model
	StudentID database.PID `json:"student_id,omitempty"`
	Student   Student      `gorm:"foreignKey:StudentID" json:"student,omitempty"`
	RoleID    database.PID `json:"role_id,omitempty"`
	Role      Role         `gorm:"foreignKey:RoleID" json:"role,omitempty"`
	FacultyID database.PID `json:"faculty_id,omitempty"`
	Faculty   DaneshKadeha `json:"faculty,omitempty"`
	Meta      string       `gorm:"type:json" json:"meta,omitempty"`
}
