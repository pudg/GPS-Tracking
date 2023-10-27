// Package database implements utulity routines for manipulating
// sqlite3 database connection.
//
// The package should only be used during the initial setup steps
// to establish database connection.

package database

import (
	"log"
	"onestep/nelson/backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Creates sqlite3 .db file and sets database global variable.
func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("onestep.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
	}

	if err := db.AutoMigrate(&models.User{}, &models.Preference{}); err != nil {
		log.Fatal("Unable to migrate models: ", err)
	}

	DB = db
}
