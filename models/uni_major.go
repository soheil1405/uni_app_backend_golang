package models

import "uni_app/database"

type UniMajor struct {
	database.Model
	DegreeLevel   string       `json:"degree_level,omitempty"`
	Duration      int          `json:"duration,omitempty"`
	Description   string       `json:"description,omitempty"`
	UniID         database.PID `json:"uni_id,omitempty"`
	Uni           Uni          `gorm:"foreignKey:UniID" json:"uni,omitempty"`
	DaneshKadehID database.PID `json:"daneshkadeh_id,omitempty"`
	DaneshKadeh   DaneshKadeh  `gorm:"foreignKey:DaneshKadehID" json:"daneshkadeh,omitempty"`
	MajorID       database.PID `json:"major_id,omitempty"`
	Major         Major        `gorm:"foreignKey:MajorID" json:"major,omitempty"`
}
