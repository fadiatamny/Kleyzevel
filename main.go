package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	PORT := 3000
	server := gin.Default()

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	server.Run(fmt.Sprintf(":%d", PORT))
}
