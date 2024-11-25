package models

import (
    "gorm.io/gorm"
)

type PlaceType struct {
    gorm.Model
    Type        string `gorm:"unique;not null"`
    Description string `gorm:"type:text"`
}