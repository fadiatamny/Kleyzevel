package utils

import (
	"fmt"
	"os"

	models "Tuneless-Treasures/models"

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

	// for local dev with no db
	// db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		return nil, err
	}

	return db, nil
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Category{}, &models.Instrument{}, &models.WorkOrder{}, &models.Worker{}, &models.Order{}, &models.Customer{})
}
