package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"onestep/nelson/backend/database"
	"onestep/nelson/backend/handlers"
	"onestep/nelson/backend/middleware"
	"onestep/nelson/backend/models"
	"onestep/nelson/backend/routes"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestEmptyUserRegister(t *testing.T) {
	router := routes.InitRouter()
	middleware.InitMiddleware(router)
	database.ConnectDatabase()
	router.POST("/api/register", handlers.Register)

	req, _ := http.NewRequest("POST", "/api/register", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestValidUserRegister(t *testing.T) {
	router := routes.InitRouter()
	middleware.InitMiddleware(router)
	database.ConnectDatabase()
	router.POST("/api/register", handlers.Register)

	user := models.CreateUser{
		Email:    "test@test.com",
		Password: "test",
	}
	user_json, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(user_json))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestDuplicateUserRegister(t *testing.T) {
	router := routes.InitRouter()
	middleware.InitMiddleware(router)
	database.ConnectDatabase()
	router.POST("/api/register", handlers.Register)

	user := models.CreateUser{
		Email:    "test@test.com",
		Password: "test",
	}
	user_json, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(user_json))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestEmptyLogin(t *testing.T) {
	router := routes.InitRouter()
	middleware.InitMiddleware(router)
	database.ConnectDatabase()
	router.POST("/api/login", handlers.Login)

	req, _ := http.NewRequest("POST", "/api/login", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestInvalidLogin(t *testing.T) {
	router := routes.InitRouter()
	middleware.InitMiddleware(router)
	database.ConnectDatabase()
	router.POST("/api/login", handlers.Login)

	user := models.CreateUser{
		Email:    "test1@test.com",
		Password: "test1",
	}
	user_json, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(user_json))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestValidLogin(t *testing.T) {
	router := routes.InitRouter()
	middleware.InitMiddleware(router)
	database.ConnectDatabase()
	router.POST("/api/login", handlers.Login)

	user := models.CreateUser{
		Email:    "test@test.com",
		Password: "test",
	}
	user_json, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(user_json))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDevices(t *testing.T) {
	router := routes.InitRouter()
	middleware.InitMiddleware(router)
	database.ConnectDatabase()
	router.GET("/api/devices", handlers.Devices)
	req, _ := http.NewRequest("GET", "/api/devices", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestValidPreferencesUpdate(t *testing.T) {
	router := routes.InitRouter()
	middleware.InitMiddleware(router)
	database.ConnectDatabase()
	router.POST("/api/preferences", handlers.UpdatePreferences)

	user := models.CreateUser{
		Email:    "test@test.com",
		Password: "test",
	}
	preference := models.Preference{
		SortAsc: true,
		Devices: []string{"{id: '8askjhfkjs', hide: false, image: null}"},
	}
	user.Preference = preference

	user_json, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/api/preferences", bytes.NewBuffer(user_json))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
