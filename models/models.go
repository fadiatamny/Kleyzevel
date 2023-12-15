package models

import (
	"time"

	_ "gorm.io/gorm"
)

type DBEntity struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
