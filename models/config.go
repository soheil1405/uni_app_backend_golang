package models

import "uni_app/database"

// Config struct holds database configuration
type Config struct {
	Database   database.Database `json:"database,omitempty"`
	ApiVersion string            `json:"api_version,omitempty"`
	Port       string            `json:"port,omitempty"`
	Auth       map[string]string `json:"auth,omitempty"`
	Migrations map[string]interface{}
}
