package services

import (
	"Tuneless-Treasures/models"

	"gorm.io/gorm"
)

func ListInstruments(db *gorm.DB) ([]models.Instrument, error) {
	instruments := []models.Instrument{}
	err := db.Find(&instruments).Error
	if err != nil {
		return nil, err
	}
	return instruments, nil
}

func getInstrument(db *gorm.DB, id string) (*models.Instrument, error) {
	instrument := models.Instrument{}
	err := db.First(&instrument, id).Error
	if err != nil {
		return nil, err
	}
	return &instrument, nil
}