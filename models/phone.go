package models

import "uni_app/database"

type Phones []*Phone

type Phone struct {
	database.Model
	PolymorphicModel
	Title string `json:"title,omitempty"`
	Phone string `json:"phone,omitempty"`
}
