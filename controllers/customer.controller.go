package controllers

import (
	"Tuneless-Treasures/models"
	"Tuneless-Treasures/services"
	"strconv"

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
	customerIdStr := c.Param("id")
	customerId, err := strconv.ParseUint(customerIdStr, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid customer id",
		})
		return
	}

	customer, err := services.GetCustomer(db, uint(customerId))
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
	customerIdStr := c.Param("id")
	customerId, err := strconv.ParseUint(customerIdStr, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid customer id",
		})
		return
	}

	payload := models.CustomerPatch{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid input",
		})
		return
	}

	customer, err := services.UpdateCustomer(db, uint(customerId), &payload)
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
	customerIdStr := c.Param("id")
	customerId, err := strconv.ParseUint(customerIdStr, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid customer id",
		})
		return
	}

	err = services.DeleteCustomer(db, uint(customerId))
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(204)
}
