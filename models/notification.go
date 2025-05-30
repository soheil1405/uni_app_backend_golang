package models

import (
	"time"
	"uni_app/database"
)

// NotificationType represents the type of notification
type NotificationType string

const (
	NotificationTypePush NotificationType = "push"
	NotificationTypeSMS  NotificationType = "sms"
)

// NotificationStatus represents the status of a notification
type NotificationStatus string

const (
	NotificationStatusPending   NotificationStatus = "pending"
	NotificationStatusSent      NotificationStatus = "sent"
	NotificationStatusFailed    NotificationStatus = "failed"
	NotificationStatusDelivered NotificationStatus = "delivered"
)

// Notification represents a notification that can be sent to users or students
type Notification struct {
	database.Model
	Type          NotificationType   `json:"type" gorm:"type:varchar(20);not null"`
	Title         string             `json:"title" gorm:"type:varchar(255);not null"`
	Body          string             `json:"body" gorm:"type:text;not null"`
	Data          map[string]string  `json:"data" gorm:"type:jsonb"`
	Status        NotificationStatus `json:"status" gorm:"type:varchar(20);default:'pending'"`
	RecipientID   database.PID       `json:"recipient_id" gorm:"not null"`
	RecipientType string             `json:"recipient_type" gorm:"type:varchar(20);not null"` // user or student
	SentAt        *time.Time         `json:"sent_at"`
	DeliveredAt   *time.Time         `json:"delivered_at"`
	Error         string             `json:"error" gorm:"type:text"`
	IsActive      bool               `json:"is_active" gorm:"default:true"`
}

// FetchNotificationRequest represents the request parameters for fetching notifications
type FetchNotificationRequest struct {
	Type          NotificationType   `json:"type" query:"type"`
	Status        NotificationStatus `json:"status" query:"status"`
	RecipientID   database.PID       `json:"recipient_id" query:"recipient_id"`
	RecipientType string             `json:"recipient_type" query:"recipient_type"`
	Page          int                `json:"page" query:"page"`
	PerPage       int                `json:"per_page" query:"per_page"`
	SortBy        string             `json:"sort_by" query:"sort_by"`
	SortDesc      bool               `json:"sort_desc" query:"sort_desc"`
}

// NotificationTemplate represents a template for notifications
type NotificationTemplate struct {
	database.Model
	Name      string           `json:"name" gorm:"type:varchar(255);not null;uniqueIndex"`
	Type      NotificationType `json:"type" gorm:"type:varchar(20);not null"`
	Title     string           `json:"title" gorm:"type:varchar(255);not null"`
	Body      string           `json:"body" gorm:"type:text;not null"`
	Variables []string         `json:"variables" gorm:"type:text[]"`
	IsActive  bool             `json:"is_active" gorm:"default:true"`
}

// NotificationPreference represents user/student notification preferences
type NotificationPreference struct {
	database.Model
	OwnerID      database.PID `json:"owner_id" gorm:"not null"`
	OwnerType    string       `json:"owner_type" gorm:"type:varchar(20);not null"` // user or student
	PushEnabled  bool         `json:"push_enabled" gorm:"default:true"`
	SMSEnabled   bool         `json:"sms_enabled" gorm:"default:true"`
	EmailEnabled bool         `json:"email_enabled" gorm:"default:true"`
	IsActive     bool         `json:"is_active" gorm:"default:true"`
}
