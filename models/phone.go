package models

import "uni_app/database"

type Phones []*Phone

type Phone struct {
	database.Model
	PolymorphicModel
	Active bool   `json:"active"`
	Title  string `json:"title,omitempty"`
	Phone  string `json:"phone,omitempty"`
}

func PhoneAcceptIncludes() []string {
	return []string{}
}
