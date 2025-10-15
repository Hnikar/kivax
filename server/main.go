// server/main.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ws", func(c *gin.Context) {
		WsHandler(c.Writer, c.Request)
	})

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Сервер на Gin работает.")
	})

	router.Run(":8080")
}
