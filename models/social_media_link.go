package models

import (
	"time"

	"gorm.io/gorm"
)

type SocialMediaLink struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Polymorphic relationship fields
	LinkableID   uint   `json:"linkable_id"`
	LinkableType string `json:"linkable_type"` // "University", "Faculty", "Place", etc.

	// Social media details
	Platform string `json:"platform"` // e.g., "instagram", "twitter", "linkedin", "telegram", "website"
	URL      string `json:"url"`
	Title    string `json:"title"`     // Optional title for the link
	Order    int    `json:"order"`     // For ordering multiple links
	IsActive bool   `json:"is_active"` // To enable/disable links
}

// TableName specifies the table name for the SocialMediaLink model
func (SocialMediaLink) TableName() string {
	return "social_media_links"
}
