package models

import "uni_app/database"

type AuthRules struct {
	database.Model
	PolymorphicModel
	V0 string `json:"v0,omitempty"`
	V1 string `json:"v1,omitempty"`
	V2 string `json:"v2,omitempty"`
	V3 string `json:"v3,omitempty"`
	V4 string `json:"v4,omitempty"`
	V5 string `json:"v5,omitempty"`
	V6 string `json:"v6,omitempty"`
	V7 string `json:"v7,omitempty"`
}
