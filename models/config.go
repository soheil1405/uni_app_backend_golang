package models

import "uni_app/database"

// Config struct holds database configuration
type Config struct {
	Database database.Database `json:"database"`
}
