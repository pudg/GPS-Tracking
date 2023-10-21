package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var upgrader = websocket.Upgrader{}
var URL = "https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key="

func LoadAPIKey(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Unable to load .env file: ", err)
	}
	return os.Getenv(key)
}

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "TODO: Implement home handler...",
	})
}

func Devices(c *gin.Context) {
	apiKey := LoadAPIKey("OS_API_KEY")
	// fmt.Printf("Requesting: to %s\n", URL+apiKey)
	resp, err := http.Get(URL + apiKey)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// data := string(body)
	// fmt.Println(data)
	c.JSON(http.StatusOK, gin.H{
		"data": body,
	})
}

func DevicesOld(c *gin.Context) {
	//TODO: Change to only accept requests from localhost:3000
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade: ", err)
		return
	}

	defer conn.Close()
	apiKey := LoadAPIKey("OS_API_KEY")
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

func Preferences(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "TODO: Implement Preferences GET...",
	})
}

func CreatePreferences(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "TODO: Implement Preferences POST...",
	})
}

func UpdatePreferences(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "TODO: Implement Preferences PUT...",
	})
}

func DeletePreferences(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "TODO: Implement Preferences DELETE...",
	})
}
