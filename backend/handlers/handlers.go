package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"onestep/nelson/backend/database"
	"onestep/nelson/backend/models"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var upgrader = websocket.Upgrader{}
var URL = "https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key="

func LoadEnvKey(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Unable to load .env file: ", err)
	}
	return os.Getenv(key)
}

func Login(c *gin.Context) {
	var input models.CreateUser

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Bind: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}

	user := models.User{}
	if result := database.DB.Where("email = ?", input.Email).First(&user); result.Error != nil {
		log.Println("Invalid authentication: ", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{"data": "Invalid email or password"})
		return
	}

	if database.CheckPasswordHash(input.Password, user.Password) {
		log.Println("Successful login: ", user.Email)
		c.JSON(http.StatusOK, gin.H{"data": "Success"})
		return
	}

	log.Println("Failed login attempty: ")
	c.JSON(http.StatusOK, gin.H{"data": "Invalid email or password"})
}

func Register(c *gin.Context) {
	var input models.CreateUser
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Bind: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := database.HashPassword(input.Password)
	if err != nil {
		log.Fatal("Unable to hash password: ", err)
		return
	}

	user := models.User{
		Email:    input.Email,
		Password: hashedPassword,
	}

	if result := database.DB.Where("email = ?", user.Email).First(&user); result.Error != nil {
		log.Println("CREATING USER")
		database.DB.Create(&user)
		c.JSON(http.StatusCreated, gin.H{"data": "Success"})
		return
	}

	log.Println("USER EXISTS")
	c.JSON(http.StatusBadRequest, gin.H{"data": "Email in use."})
}

func Devices(c *gin.Context) {
	apiKey := LoadEnvKey("OS_API_KEY")
	resp, err := http.Get(URL + apiKey)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	data := string(body)
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func DevicesOld(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade: ", err)
		return
	}

	defer conn.Close()
	apiKey := LoadEnvKey("OS_API_KEY")
	fmt.Println(apiKey)
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read: ", err)
			break
		}
		log.Printf("recv: %s", msg)
		err = conn.WriteMessage(msgType, msg)
		if err != nil {
			log.Println("write: ", err)
			break
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "TODO: Implement devices handler...",
	})
}

func UpdatePreferences(c *gin.Context) {
	var input models.CreateUser
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Bind: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}
	if result := database.DB.Where("email = ?", input.Email).First(&user); result.Error != nil {
		log.Println("Invalid email: ", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{"data": "Invalid email or password"})
		return
	}

	user.Preference.SortAsc = input.Preference.SortAsc
	user.Preference.Devices = input.Preference.Devices
	database.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func Test(c *gin.Context) {
	var users []models.User

	if err := database.DB.Model(&models.User{}).Preload("Preference").Find(&users).Error; err != nil {
		log.Panicln("Finding users: ", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func Preferences(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "TODO: Implement Preferences GET...",
	})
}

func CreatePreferences(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "TODO: Implemente Preferences POST..."})
}

func DeletePreferences(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "TODO: Implement Preferences DELETE...",
	})
}
