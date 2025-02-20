package models

import "uni_app/database"

type Student struct {
	database.Model
	Name          string       `json:"name,omitempty"`
	LastName      string       `json:"last_name,omitempty"`
	StudentCode   database.PID `json:"student_code,omitempty"`
	NationalCode  database.PID `json:"national_code,omitempty"`
	Password      string       `json:"-,omitempty"`
	MajorID       database.PID `json:"major_id,omitempty"`
	DaneshKadehID database.PID `json:"daneshkadeh_id,omitempty"`
}
