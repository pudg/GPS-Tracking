// Package models defines the database table schemas.
//
// The models package should only be used when creating initial migrations
// and during handler data binding steps.

package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// A User defines the schema for the User table to be created within the
// database.
type User struct {
	gorm.Model
	Email      string     `json:"email"`
	Password   string     `json:"password"`
	Preference Preference `json:"preference" gorm:"foreignkey:UserID"`
}

// Helper Use schema used to bind incoming user information during request
// handling steps.
type CreateUser struct {
	Email      string     `json:"email" binding:"required"`
	Password   string     `json:"password" binding:"required"`
	Preference Preference `json:"preference"`
}

// A Preference defines the schema for the Preference table to be created
// within the database.
//
// The Preference table has a one-to-one association with a user. That is, a
// user has a settings Preference.
type Preference struct {
	gorm.Model
	UserID  uint           `json:"userID"`
	SortAsc bool           `json:"sort" gorm:"default:false"`
	Devices pq.StringArray `json:"devices" gorm:"type:text[]"`
}
