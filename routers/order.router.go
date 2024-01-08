package routes

import (
	"Tuneless-Treasures/controllers"

	"github.com/gin-gonic/gin"
)

func SetupOrderRouter(r *gin.Engine) {
	// todo: needs to support proper query (and or statements parsing)
	r.GET("/orders", controllers.ListOrders)
	r.GET("/order/:id", controllers.GetOrder)
	r.POST("/order/", controllers.CreateOrder)
	r.DELETE("/order/:id", controllers.DeleteOrder)
}
