package database

import (
	"time"

	"gorm.io/gorm"
)

// Model ...
type Model struct {
	ID        PID            `json:"id" gorm:"column:id;primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (m Model) IsDeleted() bool {
	return m.DeletedAt.Valid
}
