package main

import (
	"onestep/nelson/backend/database"
	"onestep/nelson/backend/middleware"
	"onestep/nelson/backend/routes"
)

// main creates router, connects to database, initializes router middlware,
// router endpoints, and serves on localhost:8000
func main() {
	router := routes.InitRouter()
	database.ConnectDatabase()
	middleware.InitMiddleware(router)
	routes.InitRoutes(router)
	router.Run(":8000")
}
