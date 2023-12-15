package models

import (
	_ "github.com/jinzhu/gorm"
)

type Worker struct {
	DBEntity
	Name string    `json:"name" gorm:"column:name"`
	Surname string    `json:"surname" gorm:"column:surname"`
	Email string    `json:"email" gorm:"column:email"`
	Expertise string    `json:"expertise" gorm:"column:expertise"`
	Category Category `json:"category" gorm:"foreignkey:ID"`
}