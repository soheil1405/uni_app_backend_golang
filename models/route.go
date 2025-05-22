package models

import "uni_app/database"

type Route struct {
	database.Model
	Url          string
	Method       string
	RouteGroupID database.PID `json:"route_group_id,omitempty"`
	RouteGroup   RouteGroup   `gorm:"foreignKey:RouteGroupID" json:"route_group,omitempty"`
}

type FetchRouteRequest struct {
	FetchRequest
	RouteGroupID database.PID `json:"route_group_id" query:"route_group_id"`
}

func RouteAcceptIncludes() []string {
	return []string{
		"RouteGroup",
	}
}
