package models

import (
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
