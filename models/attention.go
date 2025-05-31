package models

import (
	"time"
	"uni_app/database"
)

// AttentionType represents the type of attention/alert
type AttentionType string

const (
	AttentionTypeWarning AttentionType = "warning"
	AttentionTypeError   AttentionType = "error"
	AttentionTypeInfo    AttentionType = "info"
	AttentionTypeSuccess AttentionType = "success"
)

// AttentionStatus represents the status of an attention/alert
type AttentionStatus string

const (
	AttentionStatusActive   AttentionStatus = "active"
	AttentionStatusRead     AttentionStatus = "read"
	AttentionStatusArchived AttentionStatus = "archived"
)

// Attention represents an attention/alert that needs to be shown to users or students
type Attention struct {
	database.Model
	Type          AttentionType   `json:"type" gorm:"type:varchar(20);not null"`
	Title         string          `json:"title" gorm:"type:varchar(255);not null"`
	Message       string          `json:"message" gorm:"type:text;not null"`
	Data          map[string]any  `json:"data" gorm:"type:jsonb"`
	Status        AttentionStatus `json:"status" gorm:"type:varchar(20);default:'active'"`
	RecipientID   database.PID    `json:"recipient_id" gorm:"not null"`
	RecipientType string          `json:"recipient_type" gorm:"type:varchar(20);not null"` // user or student
	ReadAt        *time.Time      `json:"read_at"`
	ArchivedAt    *time.Time      `json:"archived_at"`
	IsActive      bool            `json:"is_active" gorm:"default:true"`
}

// FetchAttentionRequest represents the request parameters for fetching attentions
type FetchAttentionRequest struct {
	Type          AttentionType   `json:"type" query:"type"`
	Status        AttentionStatus `json:"status" query:"status"`
	RecipientID   database.PID    `json:"recipient_id" query:"recipient_id"`
	RecipientType string          `json:"recipient_type" query:"recipient_type"`
	Page          int             `json:"page" query:"page"`
	PerPage       int             `json:"per_page" query:"per_page"`
	SortBy        string          `json:"sort_by" query:"sort_by"`
	SortDesc      bool            `json:"sort_desc" query:"sort_desc"`
}
