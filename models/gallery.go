package models

import (
	"uni_app/database"
)

type Gallery struct {
	database.Model
	OwnerID     uint   `json:"owner_id"`
	OwnerType   string `json:"owner_type"` // "University", "Faculty", "Place"
	ImageURL    string `json:"image_url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Order       int    `json:"order"`   // For ordering images in gallery
	IsMain      bool   `json:"is_main"` // To mark main/featured images
}

// TableName specifies the table name for the Gallery model
func (Gallery) TableName() string {
	return "galleries"
}
