package models

import (
    "gorm.io/gorm"
)

type MajorsChart struct {
    gorm.Model
    Name       string `gorm:"not null"`
    UniMajorID uint
    UniMajor   UniMajor `gorm:"foreignKey:UniMajorID"`
}