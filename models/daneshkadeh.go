package models

import "uni_app/database"

type DaneshKadeha []*DaneshKadeha

// دانشکده
type DaneshKadeh struct {
	database.Model
	Name              string          `json:"name,omitempty"`
	Description       string          `json:"description,omitempty"`
	University        Uni             `gorm:"foreignKey:UniversityID" json:"university,omitempty"`
	UniversityID      database.PID    `json:"university_id,omitempty"`
	Phones            Phones          `json:"phones,omitempty"`
	Address           Address         `json:"address,omitempty"`
	ContactWays       ContactWays     `json:"contact_ways,omitempty"`
	DaneshKadehTypeID database.PID    `json:"dan]esh_kadeh_type_id,omitempty"`
	DaneshKadehType   DaneshKadehType `json:"danesh_kadeh_type,omitempty" gorm:"foreignKey:DaneshKadehTypeID"`
}

type DaneshKadehType struct {
	database.Model
	Type        string `gorm:"unique;not null"`
	Description string `gorm:"type:text"`
}
