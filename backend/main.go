package main

import (
	"log"

	"github.com/jeypc/go-crud/config"
	"github.com/jeypc/go-crud/models"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.Item{})

	log.Println("Database Migration Berhasil!")

	log.Println("Aplikasi berjalan di port 8080")
}
