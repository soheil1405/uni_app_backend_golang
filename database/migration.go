package database

import (
	"uni_app/models"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.FailedJob{}, &models.PersonalAccessToken{}, &models.City{}, &models.PlaceType{}, &models.Place{}, &models.UniType{}, &models.Uni{},
		&models.Faculty{},
		&models.Major{},
		&models.UniMajor{},
		&models.UserRole{},
		&models.MajorsChart{},
		&models.Role{})
}
