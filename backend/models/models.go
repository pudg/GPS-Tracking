package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email      string     `json:"email"`
	Password   string     `json:"password"`
	Preference Preference `json:"preference" gorm:"foreignkey:UserID"`
}

type CreateUser struct {
	Email      string     `json:"email" binding:"required"`
	Password   string     `json:"password" binding:"required"`
	Preference Preference `json:"preference"`
}

type Preference struct {
	gorm.Model
	UserID  uint           `json:"userID"`
	SortAsc bool           `json:"sort" gorm:"default:false"`
	Devices pq.StringArray `json:"devices" gorm:"type:text[]"`
}
