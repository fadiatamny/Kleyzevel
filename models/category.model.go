package models

import (
	_ "gorm.io/gorm"
)

type Category struct {
	DBEntity
	Name string `json:"name" gorm:"column:name"`
}
