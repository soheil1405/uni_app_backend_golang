package models

import (
	"uni_app/database"
)

type SocialMediaLink struct {
	database.Model
	PolymorphicModel
	Platform string `json:"platform"` // e.g., "instagram", "twitter", "linkedin", "telegram", "website"
	URL      string `json:"url"`
	Title    string `json:"title"`     // Optional title for the link
	IsActive bool   `json:"is_active"` // To enable/disable links
}

// TableName specifies the table name for the SocialMediaLink model
func (SocialMediaLink) TableName() string {
	return "social_media_links"
}
