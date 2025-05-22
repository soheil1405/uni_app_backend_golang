package models

import "uni_app/database"

type Students []*Student
type StudentStatus int

const (
	StudentStatusActive StudentStatus = iota
	StudentStatusInactive
	StudentStatusPending
)

type Student struct {
	database.Model
	Name         string        `json:"name,omitempty"`
	LastName     string        `json:"last_name,omitempty"`
	StudentCode  database.PID  `json:"student_code,omitempty"`
	NationalCode database.PID  `json:"national_code,omitempty"`
	Password     string        `json:"-,omitempty"`
	Status       StudentStatus `json:"status,omitempty" gorm:"default:1"`

	UniID         database.PID `json:"uni_id,omitempty"`
	Uni           Uni          `gorm:"foreignKey:UniID" json:"uni,omitempty"`
	MajorID       database.PID `json:"major_id,omitempty"`
	Major         Major        `json:"major,omitempty" gorm:"foreignKey:MajorID"`
	DaneshKadehID database.PID `json:"daneshkadeh_id,omitempty"`
	DaneshKadeh   DaneshKadeh  `json:"danesh_kadeh,omitempty" gorm:"foreignKey:DaneshKadehID"`
}

type FetchStudentRequest struct {
	FetchRequest
	UniID         database.PID `json:"uni_id" query:"uni_id"`
	MajorID       database.PID `json:"major_id" query:"major_id"`
	DaneshKadehID database.PID `json:"daneshkadeh_id" query:"daneshkadeh_id"`
}

func StudentAcceptIncludes() []string {
	return []string{
		"Uni",
		"Major",
		"DaneshKadeh",
	}
}
