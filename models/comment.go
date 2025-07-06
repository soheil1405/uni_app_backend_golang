package models

import (
	"uni_app/database"
)

type CommentStatus string

var (
	CommentStatusPending  CommentStatus = "pending"
	CommentStatusActive   CommentStatus = "active"
	CommentStatusRejected CommentStatus = "rejected"
)

// Comment represents a comment that can be attached to any entity
type Comment struct {
	database.Model
	PolymorphicModel `json:"polymorphic_model,omitempty"`
	Content          string        `json:"content,omitempty" gorm:"type:text;not null"`
	UserID           database.PID  `json:"user_id,omitempty" gorm:"not null"`
	User             *User         `json:"user,omitempty" gorm:"foreignKey:UserID"`
	ParentID         *database.PID `json:"parent_id,omitempty" gorm:"default:null"` // For replies to comments
	Parent           *Comment      `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Replies          []Comment     `json:"replies,omitempty" gorm:"foreignKey:ParentID"`
	Status           CommentStatus `json:"status,omitempty"`
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
