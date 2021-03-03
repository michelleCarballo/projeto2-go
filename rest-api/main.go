package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("welcome", func(c *gin.Context) {
		firstname := c.Query("firstname")

		if firstname == "" {
			c.String(403, "Você precisa de um nome")
		} else {
			c.String(http.StatusOK, "Bem vindo "+firstname)
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
