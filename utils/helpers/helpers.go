package helpers

import (
	"strconv"

	"gorm.io/gorm"
)

// ParseUint parses a string to uint
func ParseUint(s string) uint {
	i, _ := strconv.ParseUint(s, 10, 64)
	return uint(i)
}

// PaginateTemplate is a template for pagination
type PaginateTemplate struct {
	Page    int   `json:"page"`
	PerPage int   `json:"per_page"`
	Total   int64 `json:"total"`
}

// NewPaginateTemplate creates a new pagination template
func NewPaginateTemplate(page, perPage int) *PaginateTemplate {
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 10
	}
	return &PaginateTemplate{
		Page:    page,
		PerPage: perPage,
	}
}

// Paginate applies pagination to a query
func (p *PaginateTemplate) Paginate(query *gorm.DB) *gorm.DB {
	offset := (p.Page - 1) * p.PerPage
	return query.Offset(offset).Limit(p.PerPage)
}
