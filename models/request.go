package models

import (
	"uni_app/database"

	"gorm.io/gorm"
)

type FetchRequest struct {
	ID        database.PID   `json:"id" query:"id"`
	IDs       []database.PID `json:"ids" query:"ids"`
	SearchKey string         `json:"search_key" query:"search_key"`
	Search    string         `json:"search" query:"search"`
	Limit     int            `json:"limit" query:"limit"`
	Offset    int            `json:"offset" query:"offset"`
	Page      int            `json:"page" query:"page"`
	Includes  []string       `json:"includes" query:"includes"`
	Sorts     []string       `json:"sorts" query:"sorts"`
	Filters   map[string]interface{} `json:"filters" query:"filters"`
}
func (request *FetchRequest) PrepareQuery(query *gorm.DB) *gorm.DB {
	var (
		search = request.Search
		ids    = request.IDs
		sorts  = request.Sorts
	)
	if request.Filters != nil {
		for key, value := range request.Filters {
			query = query.Where(key, value)
		}
	}

	if request.ID.IsValid() {
		ids = append(ids, request.ID)
	}

	if len(ids) > 0 {
		query = query.Where("id in (?)", ids)
	}

	if search != "" && request.SearchKey != "" {
		searchStr := "%" + search + "%"
		query = query.Where("? like ? ", request.SearchKey, searchStr)
	}

	for _, sort := range sorts {
		query = query.Order(sort)
	}

	return query
}
