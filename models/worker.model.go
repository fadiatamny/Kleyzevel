package models

import (
	_ "gorm.io/gorm"
)

type Worker struct {
	DBEntity
	Name      string      `json:"name" gorm:"column:name"`
	Surname   string      `json:"surname" gorm:"column:surname"`
	Email     string      `json:"email" gorm:"column:email"`
	Expertise []*Category `json:"category" gorm:"many2many:worker_categories"`
}
