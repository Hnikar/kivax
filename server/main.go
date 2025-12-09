package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const charset = "abcdefghijkmnpqrstuvwxyz23456789"

func generateRoomID(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	hub := NewHub()
	go hub.Run()

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // Разрешаем фронтенд
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	router.Use(cors.New(config))

	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		ServeWs(hub, c.Writer, c.Request, roomId)
	})

	router.POST("/api/rooms", func(c *gin.Context) {
		roomId := generateRoomID(6)
		c.JSON(http.StatusOK, gin.H{"roomId": roomId})
	})

	router.Run(":8080")
}
