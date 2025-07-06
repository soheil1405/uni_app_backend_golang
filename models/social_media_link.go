package models

import (
	"uni_app/database"
)

type SocialMediaLink struct {
	database.Model
	PolymorphicModel
	Platform string `json:"platform"`
	URL      string `json:"url"`
	Title    string `json:"title"` 
	Active   bool   `json:"active"`
}

// TableName specifies the table name for the SocialMediaLink model
func (SocialMediaLink) TableName() string {
	return "social_media_links"
}
