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
	BuildingID     database.PID     `gorm:"not null" json:"building_id"`
	Building       *Building        `json:"building,omitempty"`
	Name           string           `gorm:"not null" json:"name"`
	Capacity       int              `gorm:"not null" json:"capacity"`
	ClassSchedules []*ClassSchedule `json:"class_schedules,omitempty"`
}
