package models

import (
	_ "gorm.io/gorm"
)

type Order struct {
	DBEntity
	CustomerId uint     `json:"customer_id" gorm:"column:customer_id"`
	Customer   Customer `json:"customer" gorm:"foreignkey:ID"`
	Items      []string `json:"items" gorm:"type:text[]"`
	Total      float64  `json:"total" gorm:"column:total"`
	Status     string   `json:"status" gorm:"column:status"`
}
