package models

import "uni_app/database"

type MajorsCharts []*MajorsChart

// چارت رشته تحصیلی
type MajorsChart struct {
	database.Model
	Name            string       `gorm:"not null" json:"name,omitempty"`
	MajorID         database.PID `json:"major_id,omitempty"`
	Major           Major        `json:"major,omitempty"`
	TotalCountRatio float64      `gorm:"not null" json:"total_count_ratio,omitempty"`
}
type MajorChartRequest struct {
	FetchRequest
	MajorID       database.PID `json:"major_id" query:"major_id"`
	DaneshKadehID database.PID `json:"danesh_kadeh_id" query:"danesh_kadeh_id"`
	UniID         database.PID `json:"uni_id" query:"uni_id"`
}

func MajorChartAcceptIncludes() []string {
	return []string{
		"UniMajor",
		"DaneshKadeh",
		"Major",
		"Uni",
	}
}
