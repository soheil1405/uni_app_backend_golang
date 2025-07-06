package models

import "uni_app/database"

type ContactWays []*ContactWay

// Website         *string
// Email           *string
// PhoneNumber1    *string
// PhoneNumber2    *string

type ContactWay struct {
	database.Model
	PolymorphicModel
	Active  bool   `json:"active"`
	Type    string `json:"type,omitempty"`
	Content string `json:"content,omitempty"`
}

func ContactWayAcceptIncludes() []string {
	return []string{}
}
