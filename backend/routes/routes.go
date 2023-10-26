package routes

import (
	"io"
	"log"
	"onestep/nelson/backend/handlers"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

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

func InitRouter() *gin.Engine {
	InitFileLogging()
	router := gin.New()
	return router
}

func InitRoutes(router *gin.Engine) {
	router.POST("/api/login", handlers.Login)
	router.POST("/api/register", handlers.Register)
	router.GET("/api/devices", handlers.Devices)
	router.PUT("/api/preferences", handlers.UpdatePreferences)
	router.GET("/api/test", handlers.Test)
}
