package model

import (
	"time"

	"gorm.io/gorm"
)

type Gallery struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Polymorphic relationship fields
	ImageableID   uint   `json:"imageable_id"`
	ImageableType string `json:"imageable_type"` // "University", "Faculty", "Place"

	// Image details
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
