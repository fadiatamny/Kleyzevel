package models

import (
	"time"

	_ "github.com/jinzhu/gorm"
)

// Custom Gorm Model instead of
// gorm.Model

type DBEntity struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}
