package models

import "uni_app/database"

type RouteGroup struct {
	database.PID
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}
