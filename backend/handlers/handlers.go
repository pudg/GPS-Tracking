// Package handlers implementes routines for processing incoming web requests.
//
// The handlers package should only be used when mapping route paths to
// a respective function handler.

package handlers

import (
	"io"
	"log"
	"net/http"
	"onestep/nelson/backend/database"
	"onestep/nelson/backend/models"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// URL for OneStep API endpoint.
var URL = "https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key="

// Maps status codes to their corresponding http status name.
var statusCodeMap = map[int]int{
	200: http.StatusOK,
	201: http.StatusCreated,
	302: http.StatusFound,
	400: http.StatusBadRequest,
	401: http.StatusUnauthorized,
	404: http.StatusNotFound,
	500: http.StatusInternalServerError,
}

// A Result serves as the communication interface for a handler coroutine channel.
type Result struct {
	Data       gin.H
	StatusCode int
}

// LoadEnvKey loads the project environment variable file, and returns the
// specified variable.
func LoadEnvKey(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Unable to load .env file: ", err)
	}
	return os.Getenv(key)
}

// HashPassword encrypts and returns the hashed version of the specified password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash checks for password equality between the given un-encrypted
// and hashed passwords.
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Login authenticates the user if a matching record exists in the database.
//
// On Success, Login returns status code of 200 and status text OK.
//
// On Bad Request, Login returns a status code of 400 and status text BadRequest.
//
// On Invalid Credentials, Login returns a status code of 404 and status text Not Found
func Login(c *gin.Context) {
	ch := make(chan Result)
	go func(ctx *gin.Context) {
		var input models.User
		r := Result{}
		if err := ctx.ShouldBindJSON(&input); err != nil {
			log.Println("Error binding: ", err)
			r.Data = gin.H{"data": http.StatusText(http.StatusBadRequest)}
			r.StatusCode = 400
			ch <- r
			return
		}
		user := models.User{}
		if err := database.DB.Where("email = ?", input.Email).First(&user); err.Error != nil {
			r.Data = gin.H{"data": http.StatusText(http.StatusNotFound)}
			r.StatusCode = 404
			ch <- r
			return
		}
		if CheckPasswordHash(input.Password, user.Password) {
			r.Data = gin.H{"data": http.StatusText(http.StatusOK)}
			r.StatusCode = 200
			ch <- r
			return
		}
		r.Data = gin.H{"data": "Invalid Email or Password."}
		r.StatusCode = 400
		ch <- r
	}(c.Copy())

	result := <-ch
	c.JSON(statusCodeMap[result.StatusCode], result.Data)
}

// Register creates a new User record with the specified information.
//
// On Success, Register returns a status code of 200 and status text OK.
//
// On Bad Request, Register returns a status code of 400 and status text BadRequest.
//
// On Not Found, Register returns a status code of 404 and status text Not Found.
//
// On Internal Error, Register returns a status code of 500 and status text InternalServerError.
func Register(c *gin.Context) {
	ch := make(chan Result)
	go func(ctx *gin.Context) {
		var input models.CreateUser
		r := Result{}
		if err := ctx.ShouldBindJSON(&input); err != nil {
			log.Println("Error Binding: ", err)
			r.Data = gin.H{"data": http.StatusText(http.StatusBadRequest)}
			r.StatusCode = 400
			ch <- r
			return
		}

		user := models.User{}
		if err := database.DB.Where("email = ?", input.Email).First(&user); err.Error != nil {
			hashedPassword, err := HashPassword(input.Password)
			if err != nil {
				log.Fatal("Unable to hash password: ", err)
				r.Data = gin.H{"data": http.StatusText(http.StatusInternalServerError)}
				r.StatusCode = 500
				ch <- r
				return
			}
			user.Email = input.Email
			user.Password = hashedPassword
			database.DB.Create(&user)
			r.Data = gin.H{"data": http.StatusText(http.StatusCreated)}
			r.StatusCode = 201
			ch <- r
			return
		}
		r.Data = gin.H{"data": http.StatusText(http.StatusBadRequest)}
		r.StatusCode = 400
		ch <- r
	}(c.Copy())

	result := <-ch
	c.JSON(statusCodeMap[result.StatusCode], result.Data)
}

// Devices calls the OneStep API and returns a list of tracking devices.
//
// On Success, Devices returns a status code of 200 and status text of OK.
//
// On Internal Error, Devices returns a status code of 500 and status text of InternalServerError.
func Devices(c *gin.Context) {
	ch := make(chan Result)
	go func(ctx *gin.Context) {
		apiKey := LoadEnvKey("OS_API_KEY")
		resp, err := http.Get(URL + apiKey)
		r := Result{}
		if err != nil {
			//Error completing the external API call.
			r.Data = gin.H{"data": http.StatusText(http.StatusInternalServerError)}
			r.StatusCode = 500
			ch <- r
			return
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			//Error reading the API response body.
			r.Data = gin.H{"data": http.StatusText(http.StatusInternalServerError)}
			r.StatusCode = 500
			ch <- r
			return
		}
		r.Data = gin.H{"data": string(body)}
		r.StatusCode = 200
		ch <- r
	}(c.Copy())

	result := <-ch
	c.JSON(statusCodeMap[result.StatusCode], result.Data)
}

// UpdatePreferences saves user tracking settings, and uploaded thumbnails to the database.
//
// On Success, UpdatePreferences returns a status code of 200 and status text of OK.
//
// On Bad Request, UpdatePreferences returns a status code of 400 and status text of BadRequest.
//
// On Not Found, UpdatePreferences returns a status code of 404 and status text of Not Found.
func UpdatePreferences(c *gin.Context) {
	ch := make(chan Result)
	go func(ctx *gin.Context) {
		var input models.CreateUser
		r := Result{}
		if err := ctx.ShouldBindJSON(&input); err != nil {
			log.Println("Bind: ", err.Error())
			r.Data = gin.H{"data": http.StatusText(http.StatusBadRequest)}
			r.StatusCode = 400
			ch <- r
			return
		}

		user := models.User{}
		if err := database.DB.Where("email = ?", input.Email).First(&user); err.Error != nil {
			log.Println("Invalid Email: ", err.Error)
			r.Data = gin.H{"data": http.StatusText(http.StatusNotFound)}
			r.StatusCode = 404
			ch <- r
			return
		}

		user.Preference.SortAsc = input.Preference.SortAsc
		user.Preference.Devices = input.Preference.Devices
		database.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)
		r.Data = gin.H{"data": input}
		r.StatusCode = 200
		ch <- r
	}(c.Copy())

	result := <-ch
	c.JSON(statusCodeMap[result.StatusCode], result.Data)
}

// ViewDatabase is only for debugging/transparency purposes.
//
// ViewDatabase returns a list of existing user records within the database.
func ViewDatabase(c *gin.Context) {
	ch := make(chan Result)
	go func(ctx *gin.Context) {
		var users []models.User
		r := Result{}
		if err := database.DB.Model(&models.User{}).Preload("Preference").Find(&users).Error; err != nil {
			log.Println("Finding users: ", err.Error())
			r.Data = gin.H{"data": http.StatusText(http.StatusInternalServerError)}
			r.StatusCode = 500
			ch <- r
			return
		}

		r.Data = gin.H{"data": users}
		r.StatusCode = 200
		ch <- r
	}(c.Copy())
	result := <-ch
	c.JSON(statusCodeMap[result.StatusCode], result.Data)
}
