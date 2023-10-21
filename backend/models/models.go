package models

import (
	"gorm.io/gorm"
)

type Preference struct {
	gorm.Model
	Name string `json:"name"`
	Sort string `json:"sort"`
}
