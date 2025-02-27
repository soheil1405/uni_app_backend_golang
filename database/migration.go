package database

import (
	"log"

	"gorm.io/gorm"
)

func DoMigrate(db *gorm.DB, models []interface{}) {
	for _, m := range models {
		if err := db.AutoMigrate(m); err != nil {
			log.Fatal(err)
		}
	}
}
