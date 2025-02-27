package models

import (
	"time"
	"uni_app/database"
)

type Token struct {
	database.Model
	PolymorphicModel
	TokenKey   string    `json:"key"`
	Revoked    bool      `json:"revoked" gorm:"default:false"`
	ExpireTime time.Time `json:"expire_time"`
	Type       string    `json:"type"`
}
