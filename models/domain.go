package models

import "uni_app/database"

type Domain struct {
	database.Model
	Domain string       `json:"domain,omitempty"`
	UniID  database.PID `json:"uni_id,omitempty"`
	Active bool         `json:"active"`
}
