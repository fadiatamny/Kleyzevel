package models

import (
	"errors"

	_ "gorm.io/gorm"
)

type Order struct {
	DBEntity
	CustomerId string   `json:"customerId" gorm:"column:customer_id"`
	Customer   Customer `json:"customer" gorm:"foreignkey:Email"`
	Items      []string `json:"items" gorm:"type:text[]"`
	Total      float64  `json:"total" gorm:"column:total"`
	Status     string   `json:"status" gorm:"column:status"`
}

type OrderInput struct {
	CustomerId string   `json:"customerId"  binding:"required"`
	Items      []string `json:"items"`
	Total      float64  `json:"total"`
}

var ErrOrderWasCreatedWithMissingWorkOrders = errors.New("Order was created with missing work orders")