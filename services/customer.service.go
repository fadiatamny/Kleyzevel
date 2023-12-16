package services

import (
	"Tuneless-Treasures/models"

	"gorm.io/gorm"
)

func ListCustomers(db *gorm.DB) ([]models.Customer, error) {
	customers := []models.Customer{}
	err := db.Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func QueryCustomer(db *gorm.DB, payload *models.CustomerPatch) ([]models.Customer, error) {
	customers := []models.Customer{}
	err := db.Where(payload).Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func GetCustomer(db *gorm.DB, customerId uint) (*models.Customer, error) {
	customer := models.Customer{}
	err := db.First(&customer, customerId).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func CreateCustomer(db *gorm.DB, payload *models.CustomerInput) (*models.Customer, error) {
	customer := models.Customer{Email: payload.Email, Name: payload.Name, Surname: payload.Surname}
	err := db.Create(&customer).Error
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func UpdateCustomer(db *gorm.DB, customerId uint, payload *models.CustomerPatch) (*models.Customer, error) {
	customer, err := GetCustomer(db, customerId)
	if err != nil {
		return nil, err
	}

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	err = db.Model(&customer).Updates(payload).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func DeleteCustomer(db *gorm.DB, customerId uint) error {
	customer, err := GetCustomer(db, customerId)
	if err != nil {
		return err
	}

	err = db.Delete(&customer).Error
	if err != nil {
		return err
	}

	return nil
}
