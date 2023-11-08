// Package middleware implements initialization routines for gin framework middleware.
//
// The middleware package should only be used during the configuration of backend
// allowed origins and operations.

package middleware

import (
	"log"
	"onestep/nelson/backend/database"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// LoadEnvKey loads the project environment variable file, and returns the
// specified variable.
func LoadEnvKey(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Unable to load .env file: ", err)
	}
	return os.Getenv(key)
}

// InitMiddleWare adds Logging, Recovery, and CORS configuration to the specified router.
func InitMiddleware(router *gin.Engine) {
	InitSessionManagement(router)
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))
}

func InitSessionManagement(router *gin.Engine) {
	cookieSecret := LoadEnvKey("COOKIE_SECRET_KEY")
	store := gormsessions.NewStore(database.DB, true, []byte(cookieSecret))
	store.Options(sessions.Options{MaxAge: 60})
	router.Use(sessions.Sessions("ossession", store))
}

func GenerateToken() (string, error) {
	key := LoadEnvKey("TOKEN_SECRET_KEY")
	token := jwt.New(jwt.SigningMethodHS256)
	return token.SignedString([]byte(key))
}
