package models

import (
    "gorm.io/gorm"
)

type Faculty struct {
    gorm.Model
    UniversityID uint
    University   Uni   `gorm:"foreignKey:UniversityID"`
    Name         string
    Description  string
    BossID       uint
    Boss         User  `gorm:"foreignKey:BossID"`
}