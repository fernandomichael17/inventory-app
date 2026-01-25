package main

import (
	"log"

	"github.com/jeypc/go-crud/config"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDatabase()

	log.Println("Aplikasi berjalan di port 8080")
}
