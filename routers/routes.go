package routes

import (
	"Tuneless-Treasures/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/", controllers.Hello)

	SetupCustomerRouter(r)
	SetupOrderRouter(r)
}
