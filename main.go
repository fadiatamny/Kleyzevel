package main

import (
	routes "Tuneless-Treasures/routers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	PORT := 3000
	server := gin.Default()

	routes.SetupRouter(server)

	server.Run(fmt.Sprintf(":%d", PORT))
}
