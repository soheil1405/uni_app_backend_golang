package models

import (
	"time"
	"uni_app/database"

	"gorm.io/gorm"
)

// Article represents a blog post or article in the system
type Article struct {
	database.Model
	Title       string     `json:"title" gorm:"size:255;not null"`
	Content     string     `json:"content" gorm:"type:text;not null"`
	Summary     string     `json:"summary" gorm:"size:500"`
	Slug        string     `json:"slug" gorm:"size:255;uniqueIndex"`
	AuthorID    uint       `json:"author_id" gorm:"not null"`
	Author      User       `json:"author" gorm:"foreignKey:AuthorID"`
	Status      string     `json:"status" gorm:"size:20;default:'draft'"` // draft, published, archived
	PublishedAt time.Time  `json:"published_at"`
	Views       int        `json:"views" gorm:"default:0"`
	Likes       int        `json:"likes" gorm:"default:0"`
	Tags        []Tag      `json:"tags" gorm:"many2many:article_tags;"`
	Categories  []Category `json:"categories" gorm:"many2many:article_categories;"`
	Comments    []Comment  `json:"comments" gorm:"polymorphic:Owner;"`
	IsActive    bool       `json:"is_active" gorm:"default:true"`
}

// FetchArticleRequest represents the parameters for fetching articles
type FetchArticleRequest struct {
	Page       int    `json:"page" query:"page"`
	PerPage    int    `json:"per_page" query:"per_page"`
	Search     string `json:"search" query:"search"`
	Status     string `json:"status" query:"status"`
	AuthorID   uint   `json:"author_id" query:"author_id"`
	TagID      uint   `json:"tag_id" query:"tag_id"`
	CategoryID uint   `json:"category_id" query:"category_id"`
	SortBy     string `json:"sort_by" query:"sort_by"`
	SortDesc   bool   `json:"sort_desc" query:"sort_desc"`
}

// Tag represents a label that can be attached to articles
type Tag struct {
	gorm.Model
	Name        string    `json:"name" gorm:"size:50;uniqueIndex"`
	Slug        string    `json:"slug" gorm:"size:50;uniqueIndex"`
	Description string    `json:"description" gorm:"size:200"`
	Articles    []Article `json:"articles" gorm:"many2many:article_tags;"`
}

// Category represents a category that can be assigned to articles
type Category struct {
	gorm.Model
	Name        string    `json:"name" gorm:"size:50;uniqueIndex"`
	Slug        string    `json:"slug" gorm:"size:50;uniqueIndex"`
	Description string    `json:"description" gorm:"size:200"`
	Articles    []Article `json:"articles" gorm:"many2many:article_categories;"`
}

// BeforeCreate is a GORM hook that runs before creating a new article
func (a *Article) BeforeCreate(tx *gorm.DB) error {
	if a.Slug == "" {
		a.Slug = generateSlug(a.Title)
	}
	return nil
}

// BeforeUpdate is a GORM hook that runs before updating an article
func (a *Article) BeforeUpdate(tx *gorm.DB) error {
	if tx.Statement.Changed("Title") && a.Slug == "" {
		a.Slug = generateSlug(a.Title)
	}
	return nil
}

// generateSlug creates a URL-friendly slug from a title
func generateSlug(title string) string {
	// TODO: Implement proper slug generation
	// This is a placeholder implementation
	return title
}
