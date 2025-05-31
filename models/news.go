package models

import (
	"time"
	"uni_app/database"
)

// NewsStatus represents the status of a news item
type NewsStatus string

const (
	NewsStatusDraft     NewsStatus = "draft"
	NewsStatusPublished NewsStatus = "published"
	NewsStatusArchived  NewsStatus = "archived"
)

// Constants for different types of news owners
const (
	NewsOwnerTypeUni         = "uni"
	NewsOwnerTypeDaneshKadeh = "daneshkadeh"
	NewsOwnerTypePlace       = "place"
)

// News represents a news item that can be owned by different entities
type News struct {
	database.Model
	Title       string         `json:"title" gorm:"type:varchar(255);not null"`
	Content     string         `json:"content" gorm:"type:text;not null"`
	Summary     string         `json:"summary" gorm:"type:varchar(500)"`
	Slug        string         `json:"slug" gorm:"type:varchar(255);uniqueIndex"`
	Status      NewsStatus     `json:"status" gorm:"type:varchar(20);default:'draft'"`
	PublishedAt *time.Time     `json:"published_at"`
	OwnerID     database.PID   `json:"owner_id" gorm:"not null"`
	OwnerType   string         `json:"owner_type" gorm:"type:varchar(20);not null"` // uni, daneshkadeh, place
	AuthorID    database.PID   `json:"author_id" gorm:"not null"`
	Author      *User          `json:"author,omitempty" gorm:"foreignKey:AuthorID"`
	Tags        []Tag          `json:"tags,omitempty" gorm:"many2many:news_tags;"`
	Views       int            `json:"views" gorm:"default:0"`
	Likes       int            `json:"likes" gorm:"default:0"`
	Comments    []Comment      `json:"comments,omitempty" gorm:"polymorphic:Owner;"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	IsNotified  bool           `json:"is_notified" gorm:"default:false"`
	NotifyAt    *time.Time     `json:"notify_at"`
	Meta        map[string]any `json:"meta" gorm:"type:jsonb"`
}

// FetchNewsRequest represents the request parameters for fetching news
type FetchNewsRequest struct {
	Status    NewsStatus   `json:"status" query:"status"`
	OwnerID   database.PID `json:"owner_id" query:"owner_id"`
	OwnerType string       `json:"owner_type" query:"owner_type"`
	AuthorID  database.PID `json:"author_id" query:"author_id"`
	TagIDs    []uint       `json:"tag_ids" query:"tag_ids"`
	Page      int          `json:"page" query:"page"`
	PerPage   int          `json:"per_page" query:"per_page"`
	SortBy    string       `json:"sort_by" query:"sort_by"`
	SortDesc  bool         `json:"sort_desc" query:"sort_desc"`
	Search    string       `json:"search" query:"search"`
}

func NewsAcceptIncludes() []string {
	return []string{
		"Author",
		"Tags",
		"Comments",
	}
}
