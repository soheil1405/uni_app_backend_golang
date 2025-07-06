package models

import "uni_app/database"

type From struct {
	database.Model
	Title       string       `json:"title,omitempty"`
	Description string       `json:"description,omitempty"`
	UniID       database.PID `json:"uni_id,omitempty"`
	Uni         Uni          `json:"uni,omitempty"`
	Link        string       `json:"link,omitempty"`
}
