// Package routes defines utility routines for creating a router, initializing
// file logs, and mapping routes to corresponding handler functions.
package routes

import (
	"io"
	"log"
	"onestep/nelson/backend/handlers"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// InitFileLogging configures gin to log all traffic to the created file, and stdout.
func InitFileLogging() {
	gin.DisableConsoleColor()
	currentTime := time.Now()
	f, err := os.Create("logs/" + currentTime.Format("Mon 2006-01-2") + ".log")
	if err != nil {
		log.Fatal("Creating .log file: ", err.Error())
		return
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

// InitRouter creates a router, and initializes file logging.
func InitRouter() *gin.Engine {
	InitFileLogging()
	router := gin.New()
	router.SetTrustedProxies(nil)
	return router
}

// InitRoutes maps API endpoints to corresponding function handlers.
func InitRoutes(router *gin.Engine) {
	router.POST("/api/login", handlers.Login)
	router.POST("/api/register", handlers.Register)
	router.GET("/api/devices", handlers.Devices)
	router.PUT("/api/preferences", handlers.UpdatePreferences)
	router.GET("/api/database", handlers.ViewDatabase)
}
