package controllers

import (
	"Tuneless-Treasures/models"
	"Tuneless-Treasures/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListOrders(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	orders, err := services.ListOrders(db)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, orders)
}

func GetOrder(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	order, err := services.GetOrder(db, id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, order)
}

func CreateOrder(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	payload := models.OrderInput{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid input",
		})
		return
	}

	order, err := services.CreateOrder(db, &payload)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, order)
}

func DeleteOrder(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	err := services.DeleteOrder(db, id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(204, nil)
}
