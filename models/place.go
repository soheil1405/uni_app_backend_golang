package models

import (
    "gorm.io/gorm"
)

type Place struct {
    gorm.Model
    Name         string     `gorm:"not null"`
    CityID       uint       `gorm:"foreignKey:ID"`
    City         City       `gorm:"foreignKey:CityID"`
    PlaceTypeID  uint       `gorm:"foreignKey:ID"`
    PlaceType    PlaceType  `gorm:"foreignKey:PlaceTypeID"`
}