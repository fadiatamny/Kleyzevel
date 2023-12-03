package utils

import (
	"fmt"
	"os"

	models "Tuneless-Treasures/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s", dbHost, dbPort, dbUser, dbName, sslMode, dbPassword))

	if err != nil {
		return nil, err
	}

	return db, nil
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Instrument{})
}

func CloseDB(db *gorm.DB) {
	db.Close()
}
