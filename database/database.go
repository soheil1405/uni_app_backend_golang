package database

import (
	"uni_app/models"

	"gorm.io/gorm"
)

func Init(db *gorm.DB) error {
	// Auto migrate models
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return err
	}

	return nil
}
