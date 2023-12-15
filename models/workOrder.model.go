package models

import (
	_ "gorm.io/gorm"
)

type WorkOrder struct {
	DBEntity
	InstrumentID uint       `json:"instrument_id" gorm:"column:instrument_id"`
	Instrument   Instrument `json:"instrument" gorm:"foreignkey:ID"`
	WorkerId     uint       `json:"worker_id" gorm:"column:worker_id"`
	Worker       Worker     `json:"worker" gorm:"foreignkey:ID"`
	Status       string     `json:"status" gorm:"column:status"`
	OrderId      uint       `json:"order_id" gorm:"column:order_id"`
	Order        Order      `json:"order" gorm:"foreignkey:ID"`
}
