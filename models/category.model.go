package models

import (
	_ "github.com/jinzhu/gorm"
)

type Category struct {
	DBEntity
	Name string `json:"name" gorm:"column:name"`
}