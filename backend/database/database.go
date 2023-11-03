// Package database implements utulity routines for manipulating
// sqlite3 database connection.
//
// The package should only be used during the initial setup steps
// to establish database connection.

package database

import (
	"log"
	"net/http"
	"onestep/nelson/backend/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// A Result serves as the communication interface for a handler coroutine channel.
type Result struct {
	Data       gin.H
	StatusCode int
}

// CheckPasswordHash checks for password equality between the given un-encrypted
// and hashed passwords.
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// HashPassword encrypts and returns the hashed version of the specified password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

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

// AuthenticateUser checks if a matching User record exists in the database.
//
// Returns a Result interface containing the operation status code and corresponding data.
func AuthenticateUser(input models.User) Result {
	user := models.User{}
	r := Result{}
	if err := DB.Where("email = ?", input.Email).First(&user); err.Error != nil {
		r.Data = gin.H{"data": http.StatusText(http.StatusNotFound)}
		r.StatusCode = 404
		return r
	}
	if CheckPasswordHash(input.Password, user.Password) {
		r.Data = gin.H{"data": http.StatusText(http.StatusOK)}
		r.StatusCode = 200
		return r
	}
	r.Data = gin.H{"data": "Invalid Email or Password."}
	r.StatusCode = 400
	return r
}

// CreateUser creates a new User record with the specified information.
//
// Returns a Result interface containing the operation status code and corresponding data.
func CreateUser(input models.CreateUser) Result {
	user := models.User{}
	r := Result{}
	if err := DB.Where("email = ?", input.Email).First(&user); err.Error != nil {
		hashedPassword, err := HashPassword(input.Password)
		if err != nil {
			log.Fatal("Unable to hash password: ", err)
			r.Data = gin.H{"data": http.StatusText(http.StatusInternalServerError)}
			r.StatusCode = 500
			return r
		}
		user.Email = input.Email
		user.Password = hashedPassword
		DB.Create(&user)
		r.Data = gin.H{"data": http.StatusText(http.StatusCreated)}
		r.StatusCode = 201
		return r
	}
	r.Data = gin.H{"data": http.StatusText(http.StatusBadRequest)}
	r.StatusCode = 400
	return r
}

// UpdateUserPreferences patches the device Preferences for existing User.
//
// Returns Result interface containing the operation status code and corresponding data.
func UpdateUserPreferences(input models.CreateUser) Result {
	r := Result{}
	user := models.User{}
	if err := DB.Where("email = ?", input.Email).First(&user); err.Error != nil {
		log.Println("Invalid Email: ", err.Error)
		r.Data = gin.H{"data": http.StatusText(http.StatusNotFound)}
		r.StatusCode = 404
		return r
	}
	user.Preference.SortAsc = input.Preference.SortAsc
	user.Preference.Devices = input.Preference.Devices
	DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)
	r.Data = gin.H{"data": input}
	r.StatusCode = 200
	return r
}

// AllUser pulls all existing user records from the database.
//
// AllUsers should only be used for testing purposes to view the existing database.
//
// Returns a Result interface containing the operation status and corresponding data.
func AllUsers() Result {
	var users []models.User
	r := Result{}
	if err := DB.Model(&models.User{}).Preload("Preference").Find(&users).Error; err != nil {
		log.Println("Finding Users: ", err.Error())
		r.Data = gin.H{"data": http.StatusText(http.StatusInternalServerError)}
		r.StatusCode = 500
		return r
	}
	r.Data = gin.H{"data": users}
	r.StatusCode = 200
	return r
}
