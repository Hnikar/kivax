package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	hub := NewHub()
	go hub.Run()

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST"}
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	router.Use(cors.New(config))

	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		ServeWs(hub, c.Writer, c.Request, roomId)
	})

	router.POST("/api/rooms", func(c *gin.Context) {

	})

	router.Run(":8080")
}
