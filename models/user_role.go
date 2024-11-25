package models

import (
    "gorm.io/gorm"
)

type UserRole struct {
    gorm.Model
    UserID uint
    User   User `gorm:"foreignKey:UserID"`
    RoleID uint
    Role   Role `gorm:"foreignKey:RoleID"`
    Meta   string `gorm:"type:json"`
}