package models

import (
	"errors"
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

type CustomerInput struct {
	Email   string `json:"email" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Surname string `json:"surname" binding:"required"`
}

type CustomerPatch struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func (c *CustomerPatch) Validate() error {
	if c.Email == "" && c.Name == "" && c.Surname == "" {
		return errors.New("at least one field must be provided")
	}
	return nil
}
