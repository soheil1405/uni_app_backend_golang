package models

import "uni_app/database"

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
