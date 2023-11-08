// Package handlers implementes routines for processing incoming web requests.
//
// The handlers package should only be used when mapping route paths to
// a respective function handler.

package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"onestep/nelson/backend/database"
	"onestep/nelson/backend/models"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

// LoadEnvKey loads the project environment variable file, and returns the
// specified variable.
func LoadEnvKey(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Unable to load .env file: ", err)
	}
	return os.Getenv(key)
}

// Login handles the validation and authentication process for login requests.
//
// On Success, Login returns status code of 200 and status text OK.
//
// On Bad Request, Login returns a status code of 400 and status text BadRequest.
//
// On Invalid Credentials, Login returns a status code of 404 and status text Not Found
func Login(c *gin.Context) {
	ch := make(chan database.Result)
	go func(ctx *gin.Context) {
		var input models.User
		r := database.Result{}
		if err := ctx.ShouldBindJSON(&input); err != nil {
			r.Data = gin.H{"data": http.StatusText(http.StatusBadRequest)}
			r.StatusCode = 400
			ch <- r
			return
		}
		r = database.AuthenticateUser(input)
		if r.StatusCode == 200 {
			session := sessions.Default(c)
			bytes, err := json.Marshal(input)
			if err != nil {
				r.Data = gin.H{"data": http.StatusText(http.StatusInternalServerError)}
				r.StatusCode = 500
				ch <- r
				return
			}
			userID := r.Data["data"]
			session.Set(userID, bytes)
			err = session.Save()
			if err != nil {
				r.Data = gin.H{"data": http.StatusText(http.StatusInternalServerError)}
				r.StatusCode = 500
				ch <- r
				return
			}
			r.Data = gin.H{"data": userID}
		}
		ch <- r
	}(c.Copy())

	result := <-ch
	c.JSON(statusCodeMap[result.StatusCode], result.Data)
}

// Logout handles the cleanup process for removing a users active session.
//
// On Success, Logout reteurns a status code of 200 and status text OK.
//
// On Internal Error, Logout returns a status code of 500, and status text InternalServerError.
func Logout(c *gin.Context) {
	ch := make(chan database.Result)
	go func() {
		r := database.Result{}
		session := sessions.Default(c)
		session.Clear()
		session.Options(sessions.Options{MaxAge: -1})
		err := session.Save()
		if err != nil {
			r.Data = gin.H{"data": http.StatusText(http.StatusInternalServerError)}
			r.StatusCode = 500
			ch <- r
			return
		}
		r.Data = gin.H{"data": http.StatusText(http.StatusOK)}
		r.StatusCode = 200
		ch <- r
	}()
	result := <-ch
	c.JSON(result.StatusCode, result.Data)
}

// Register handles the validation and authentication process for register requests.
//
// On Success, Register returns a status code of 200 and status text OK.
//
// On Bad Request, Register returns a status code of 400 and status text BadRequest.
//
// On Not Found, Register returns a status code of 404 and status text Not Found.
//
// On Internal Error, Register returns a status code of 500 and status text InternalServerError.
func Register(c *gin.Context) {
	ch := make(chan database.Result)
	go func(ctx *gin.Context) {
		var input models.CreateUser
		r := database.Result{}
		if err := ctx.ShouldBindJSON(&input); err != nil {
			r.Data = gin.H{"data": http.StatusText(http.StatusBadRequest)}
			r.StatusCode = 400
			ch <- r
			return
		}
		r = database.CreateUser(input)
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
	ch := make(chan database.Result)
	go func(ctx *gin.Context) {
		r := database.Result{}
		if _, err := ctx.Cookie("ossession"); err != nil {
			r.Data = gin.H{"data": http.StatusText(http.StatusUnauthorized)}
			r.StatusCode = 401
			ch <- r
			return
		}

		apiKey := LoadEnvKey("OS_API_KEY")
		resp, err := http.Get(URL + apiKey)
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
	ch := make(chan database.Result)
	go func(ctx *gin.Context) {
		r := database.Result{}
		if _, err := ctx.Cookie("ossession"); err != nil {
			r.Data = gin.H{"data": http.StatusText(http.StatusUnauthorized)}
			r.StatusCode = 401
			ch <- r
			return
		}

		var preferences models.Preference
		if err := ctx.ShouldBindJSON(&preferences); err != nil {
			r.Data = gin.H{"data": http.StatusText(http.StatusBadRequest)}
			r.StatusCode = 400
			ch <- r
			return
		}

		user := models.User{}
		session := sessions.Default(ctx)
		userCredentials := session.Get(preferences.UserID)
		if userCredentials == nil {
			r.Data = gin.H{"data": http.StatusText(http.StatusUnauthorized)}
			r.StatusCode = 401
			ch <- r
			return
		}
		err := json.Unmarshal(userCredentials.([]byte), &user)
		if err != nil {
			r.Data = gin.H{"data": http.StatusText(http.StatusInternalServerError)}
			r.StatusCode = 500
			ch <- r
			return
		}

		var input models.CreateUser
		input.Email = user.Email
		input.Password = user.Password
		input.Preference = preferences
		r = database.UpdateUserPreferences(input)
		ch <- r
	}(c.Copy())

	result := <-ch
	c.JSON(statusCodeMap[result.StatusCode], result.Data)
}

// ViewDatabase is only for debugging/transparency purposes.
//
// ViewDatabase returns a list of existing user records within the database.
func ViewDatabase(c *gin.Context) {
	ch := make(chan database.Result)
	go func(ctx *gin.Context) {
		r := database.AllUsers()
		ch <- r
	}(c.Copy())
	result := <-ch
	c.JSON(statusCodeMap[result.StatusCode], result.Data)
}
