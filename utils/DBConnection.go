package utils

import (
	"fmt"
	"os"

	"Tuneless-Treasures/models"

	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	disableSSL := os.Getenv("DISABLE_SSL")

	sslMode := "disable"
	if disableSSL == "false" {
		sslMode = "enable"
	}

	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s", dbHost, dbPort, dbUser, dbName, sslMode, dbPassword)))

	if err != nil {
		return nil, err
	}

	return db, nil
}

func AutoMigrate(db *gorm.DB) error {

	if err := db.AutoMigrate(&models.WorkOrder{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.Worker{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.Order{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.Customer{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.Category{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.Instrument{}); err != nil {
		return err
	}

	return nil
}
