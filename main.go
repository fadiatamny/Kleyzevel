package main

import (
	routes "Tuneless-Treasures/routers"
	utils "Tuneless-Treasures/utils"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}

func main() {
	loadEnv()
	log.Println("Successfully loaded env config")

	db, connectionErr := utils.ConnectDB()
	if connectionErr != nil {
		log.Fatal("Error connecting to database", connectionErr)
	}
	log.Println("Successfully connected to database")

	migrateErr := utils.AutoMigrate(db)
	if migrateErr != nil {
		log.Fatal("Error migrating database", migrateErr)
	}

	log.Println("Successfully migrated database")

	PORT := 3000
	server := gin.Default()

	server.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	routes.SetupRouter(server)

	server.Run(fmt.Sprintf(":%d", PORT))
	log.Println(fmt.Sprintf("Server running on port %d", PORT))
}
