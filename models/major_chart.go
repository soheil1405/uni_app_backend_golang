package models

import "uni_app/database"

type MajorsCharts []*MajorsChart

// چارت رشته تحصیلی
type MajorsChart struct {
	database.Model
	Name            string       `gorm:"not null" json:"name,omitempty"`
	UniMajorID      database.PID `json:"uni_major_id,omitempty"`
	UniMajor        UniMajor     `gorm:"foreignKey:UniMajorID" json:"uni_major,omitempty"`
	DaneshKadehID   database.PID `json:"danesh_kadeh_id,omitempty"`
	DaneshKadeh     DaneshKadeh  `gorm:"foreignKey:DaneshKadehID" json:"danesh_kadeh,omitempty"`
	MajorID         uint         `gorm:"not null" json:"major_id"`
	Major           Major        `gorm:"foreignKey:MajorID" json:"major"`
	UniID           uint         `gorm:"not null" json:"uni_id"`
	Uni             Uni          `gorm:"foreignKey:UniID" json:"uni"`
	TotalCountRatio float64      `gorm:"not null" json:"total_count_ratio"`
}
