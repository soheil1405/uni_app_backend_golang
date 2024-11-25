package models

import (
    "gorm.io/gorm"
    "time"
)

type Uni struct {
    gorm.Model
    Name           string
    ZipCode        *int
    Website        *string
    Email          *string
    PhoneNumber1   *string
    PhoneNumber2   *string
    UniTypeID      uint
    UniType        UniType `gorm:"foreignKey:UniTypeID"`
    BossID         uint
    Boss           User    `gorm:"foreignKey:BossID"`
    CityID         uint
    City           City    `gorm:"foreignKey:CityID"`
    EstablishedYear *time.Time
}