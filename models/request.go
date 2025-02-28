package models

import (
	"fmt"
	"uni_app/database"

	"gorm.io/gorm"
)

type FetchRequest struct {
	ID       database.PID           `json:"id" query:"id"`
	IDs      []database.PID         `json:"ids" query:"ids"`
	Search   string                 `json:"search" query:"search"`
	Limit    int                    `json:"limit" query:"limit"`
	Offset   int                    `json:"offset" query:"offset"`
	Page     int                    `json:"page" query:"page"`
	Filters  map[string]interface{} `json:"filters" query:"filters"`
	Includes []string               `json:"includes" query:"includes"`
	Sorts    []string               `json:"sorts" query:"sorts"`
}

func (request *FetchRequest) PrepareQuery(query *gorm.DB) *gorm.DB {
	var (
		search  = request.Search
		ids     = request.IDs
		filters = request.Filters
		sorts   = request.Sorts
	)

	for key, value := range filters {
		switch v := value.(type) {
		case string:
			query = query.Where(fmt.Sprintf("%s = ?", key), v)
		case []string:
			query = query.Where(fmt.Sprintf("%s IN (?)", key), v)
		default:
			fmt.Printf("Unexpected type for key %s: %T\n", key, v)
		}
	}

	if len(ids) > 0 {
		query = query.Where("id in (?)", ids)
	}

	if search != "" {
		searchStr := "%" + search + "%"
		query = query.Where("name like ? ", searchStr)
	}

	for _, sort := range sorts {
		query = query.Order(sort)
	}

	return query
}
