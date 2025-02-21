package models

type PolymorphicModel struct {
	OwnerType string `json:"owner_type,omitempty"`
	OwnerID   string `json:"owner_id,omitempty"`
}
