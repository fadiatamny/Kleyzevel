package routes

import (
	"Tuneless-Treasures/controllers"

	"github.com/gin-gonic/gin"
)

func SetupCustomerRouter(r *gin.Engine) {
	// todo: needs to support proper query (and or statements parsing)
	r.GET("/customers", controllers.ListCustomers)
	r.POST("/customers/query", controllers.QueryCustomer)
	r.GET("/customer/:id", controllers.GetCustomer)
	r.POST("/customer/", controllers.CreateCustomer)
	r.PATCH("/customer/:id", controllers.UpdateCustomer)
	r.DELETE("/customer/:id", controllers.DeleteCustomer)
}
