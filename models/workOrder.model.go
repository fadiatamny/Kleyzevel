package models

import (
	_ "gorm.io/gorm"
)

type WorkOrder struct {
	DBEntity
	InstrumentID uint       `json:"instrumentId" gorm:"column:instrument_id"`
	Instrument   Instrument `json:"instrument" gorm:"foreignkey:ID"`
	WorkerId     uint       `json:"workerId" gorm:"column:worker_id"`
	Worker       Worker     `json:"worker" gorm:"foreignkey:ID"`
	Status       string     `json:"status" gorm:"column:status"`
	OrderId      uint       `json:"orderId" gorm:"column:order_id"`
	Order        Order      `json:"order" gorm:"foreignkey:ID"`
}
