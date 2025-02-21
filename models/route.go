package models

import "uni_app/database"

type Route struct {
	database.Model
	Url          string
	Method       string
	RouteGroupID database.PID
	RouteGroup   RouteGroup
}
