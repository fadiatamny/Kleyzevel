package controllers

import (
	"Tuneless-Treasures/models"
	"Tuneless-Treasures/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListCustomers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	customers, err := services.ListCustomers(db)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, customers)
}

func QueryCustomer(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	payload := models.CustomerPatch{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid input",
		})
		return
	}

	customers, err := services.QueryCustomer(db, &payload)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, customers)
}

func GetCustomer(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	email := c.Param("email")

	customer, err := services.GetCustomer(db, email)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, customer)
}

func CreateCustomer(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	payload := models.CustomerInput{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid input",
		})
		return
	}

	customer, err := services.CreateCustomer(db, &payload)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, customer)
}

func UpdateCustomer(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	email := c.Param("email")

	payload := models.CustomerPatch{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid input",
		})
		return
	}

	customer, err := services.UpdateCustomer(db, email, &payload)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, customer)
}

func DeleteCustomer(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	email := c.Param("email")

	err := services.DeleteCustomer(db, email)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(204)
}
