package models

import (
	"uni_app/database"
)

type RoomType string

const (
	RoomTypeClassroom RoomType = "classroom"
	RoomTypeOffice    RoomType = "office"
	RoomTypeLab       RoomType = "lab"
	RoomTypeOther     RoomType = "other"
)

type FetchRoomRequest struct {
	FetchRequest
	BuildingID database.PID `json:"building_id,omitempty"`
	Name       string       `json:"name,omitempty"`
	Capacity   int          `json:"capacity,omitempty"`
}

type Room struct {
	database.Model
	FloorID         database.PID      `json:"floor_id,omitempty"`
	Floor           Floor             `json:"floor,omitempty"`
	Name            string            `gorm:"not null" json:"name,omitempty"`
	Capacity        int               `gorm:"not null" json:"capacity,omitempty"`
	Active          bool              `json:"active"`
	CourseInstances []*CourseInstance `json:"class_instances,omitempty"`
}
