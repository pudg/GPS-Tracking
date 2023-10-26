package database

import (
	"fmt"
	"log"
	"onestep/nelson/backend/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("onestep.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
	}

	if err := db.AutoMigrate(&models.User{}, &models.Preference{}); err != nil {
		log.Fatal("Unable to migrate models: ", err)
	}

	DB = db
	fmt.Println("COMPLETE: Connected to database.")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
