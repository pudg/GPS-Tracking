package models

import (
	"gorm.io/gorm"
)

type Preference struct {
	gorm.Model
	Name string `json:"name"`
	Sort string `json:"sort"`
}

type CreatePreference struct {
	Name string `json:"name"`
	Sort string `json:"sort"`
}

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
