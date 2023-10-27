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
	"gorm.io/gorm"
)

var URL = "https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key="
var statusCodeMap = map[int]int{
	200: http.StatusOK,
	201: http.StatusCreated,
	302: http.StatusFound,
	400: http.StatusBadRequest,
	401: http.StatusUnauthorized,
	404: http.StatusNotFound,
	500: http.StatusInternalServerError,
}

type Result struct {
	Data       gin.H
	StatusCode int
}

func LoadEnvKey(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Unable to load .env file: ", err)
	}
	return os.Getenv(key)
}

func Login(c *gin.Context) {
	ch := make(chan Result)
	go func(ctx *gin.Context) {
		var input models.User
		r := Result{}
		if err := ctx.ShouldBindJSON(&input); err != nil {
			log.Println("Error binding: ", err)
			r.Data = gin.H{"data": http.StatusText(400)}
			r.StatusCode = 400
			ch <- r
			return
		}
		user := models.User{}
		if err := database.DB.Where("email = ?", input.Email).First(&user); err.Error != nil {
			r.Data = gin.H{"data": http.StatusText(404)}
			r.StatusCode = 404
			ch <- r
			return
		}
		if database.CheckPasswordHash(input.Password, user.Password) {
			r.Data = gin.H{"data": http.StatusText(200)}
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

func Register(c *gin.Context) {
	ch := make(chan Result)
	go func(ctx *gin.Context) {
		var input models.CreateUser
		r := Result{}
		if err := ctx.ShouldBindJSON(&input); err != nil {
			log.Println("Error Binding: ", err)
			r.Data = gin.H{"data": http.StatusText(400)}
			r.StatusCode = 400
			ch <- r
			return
		}

		user := models.User{}
		if err := database.DB.Where("email = ?", input.Email).First(&user); err.Error != nil {
			hashedPassword, err := database.HashPassword(input.Password)
			if err != nil {
				log.Fatal("Unable to hash password: ", err)
				r.Data = gin.H{"data": http.StatusText(500)}
				r.StatusCode = 500
				ch <- r
				return
			}
			user.Email = input.Email
			user.Password = hashedPassword
			database.DB.Create(&user)
			r.Data = gin.H{"data": http.StatusText(200)}
			r.StatusCode = 200
			ch <- r
			return
		}
		r.Data = gin.H{"data": http.StatusText(400)}
		r.StatusCode = 400
		ch <- r
	}(c.Copy())

	result := <-ch
	c.JSON(statusCodeMap[result.StatusCode], result.Data)
}

func Devices(c *gin.Context) {
	ch := make(chan Result)
	go func(ctx *gin.Context) {
		apiKey := LoadEnvKey("OS_API_KEY")
		resp, err := http.Get(URL + apiKey)
		r := Result{}
		if err != nil {
			r.Data = gin.H{"data": http.StatusText(500)}
			r.StatusCode = 500
			ch <- r
			return
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			r.Data = gin.H{"Data": http.StatusText(500)}
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

func UpdatePreferences(c *gin.Context) {
	ch := make(chan Result)
	go func(ctx *gin.Context) {
		var input models.CreateUser
		r := Result{}
		if err := ctx.ShouldBindJSON(&input); err != nil {
			log.Println("Bind: ", err.Error())
			r.Data = gin.H{"data": http.StatusText(400)}
			r.StatusCode = 400
			ch <- r
			return
		}

		user := models.User{}
		if err := database.DB.Where("email = ?", input.Email).First(&user); err.Error != nil {
			log.Println("Invalid Email: ", err.Error)
			r.Data = gin.H{"data": http.StatusText(404)}
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

func ViewDatabase(c *gin.Context) {
	ch := make(chan Result)
	go func(ctx *gin.Context) {
		var users []models.User
		r := Result{}
		if err := database.DB.Model(&models.User{}).Preload("Preference").Find(&users).Error; err != nil {
			log.Println("Finding users: ", err.Error())
			r.Data = gin.H{"data": http.StatusText(500)}
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
