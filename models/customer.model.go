package models

import (
	"time"

	_ "gorm.io/gorm"
)

type Customer struct {
	Email     string    `json:"email" gorm:"column:email;primary_key"`
	Name      string    `json:"name" gorm:"column:name"`
	Surname   string    `json:"surname" gorm:"column:surname"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}
