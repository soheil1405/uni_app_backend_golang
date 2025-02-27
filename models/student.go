package models

import "uni_app/database"

type Students []*Student

type Student struct {
	database.Model
	Name          string       `json:"name,omitempty"`
	LastName      string       `json:"last_name,omitempty"`
	StudentCode   database.PID `json:"student_code,omitempty"`
	NationalCode  database.PID `json:"national_code,omitempty"`
	Password      string       `json:"-,omitempty"`
	UniID         database.PID `json:"uni_id,omitempty"`
	Uni           Uni          `gorm:"foreignKey:UniID" json:"uni,omitempty"`
	MajorID       database.PID `json:"major_id,omitempty"`
	DaneshKadehID database.PID `json:"daneshkadeh_id,omitempty"`
	Status        UserMode     `json:"status,omitempty" gorm:"default:1"`
	Roles         Roles        `gorm:"many2many:student_roles;" json:"roles,omitempty"`
}
