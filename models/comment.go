package models

import (
	"uni_app/database"
)

// Comment represents a comment that can be attached to any entity
type Comment struct {
	database.Model
	Content   string        `json:"content" gorm:"type:text;not null"`
	UserID    database.PID  `json:"user_id" gorm:"not null"`
	User      *User         `json:"user,omitempty" gorm:"foreignKey:UserID"`
	OwnerID   database.PID  `json:"owner_id" gorm:"not null"`
	OwnerType string        `json:"owner_type" gorm:"not null"`
	ParentID  *database.PID `json:"parent_id" gorm:"default:null"` // For replies to comments
	Parent    *Comment      `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Replies   []Comment     `json:"replies,omitempty" gorm:"foreignKey:ParentID"`
	IsActive  bool          `json:"is_active" gorm:"default:true"`
}

// FetchCommentRequest represents the request parameters for fetching comments
type FetchCommentRequest struct {
	UserID          database.PID  `json:"user_id" query:"user_id"`
	CommentableID   database.PID  `json:"commentable_id" query:"commentable_id"`
	CommentableType string        `json:"commentable_type" query:"commentable_type"`
	ParentID        *database.PID `json:"parent_id" query:"parent_id"`
	Includes        []string      `json:"includes" query:"includes"`
}

func CommentAcceptIncludes() []string {
	return []string{
		"User",
		"Parent",
		"Replies",
	}
}
