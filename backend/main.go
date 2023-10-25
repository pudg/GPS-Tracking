package main

import (
	"onestep/nelson/backend/database"
	"onestep/nelson/backend/middleware"
	"onestep/nelson/backend/routes"
)

func main() {
	router := routes.InitRouter()
	database.ConnectDatabase()
	middleware.InitMiddleware(router)
	routes.InitRoutes(router)
	router.Run(":8000")
}
