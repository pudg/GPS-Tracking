package routes

import (
	"onestep/nelson/backend/handlers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	return router
}

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", handlers.Home)
	router.GET("/devices", handlers.Devices)
	router.GET("/preferences/:id", handlers.Preferences)
	router.POST("/preferences", handlers.CreatePreferences)
	router.PUT("/preferences/:id", handlers.UpdatePreferences)
	router.DELETE("/preferences/:id", handlers.DeletePreferences)
}
