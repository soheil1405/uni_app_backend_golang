package models

import (
	"time"
	"uni_app/database"
)

type Token struct {
	database.Model
	TokenKey   string       `json:"key"`
	Revoked    bool         `json:"revoked" gorm:"default:'false'"`
	ExpireTime time.Time    `json:"expire_time"`
	User       *User        `json:"user" gorm:"foreignkey:user_id;association_foreignkey:id"`
	OwnerType  string       `json:"owner_type"`
	OwnerID    database.PID `json:"owner_id" gorm:"index"`
	Type       string       `json:"type"`
}
