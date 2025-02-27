package models

import "uni_app/database"

type PolymorphicModel struct {
	OwnerType string       `json:"owner_type,omitempty"`
	OwnerID   database.PID `json:"owner_id,omitempty"`
}
