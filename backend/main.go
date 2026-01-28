package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jeypc/go-crud/config"
	"github.com/jeypc/go-crud/controllers"
	"github.com/jeypc/go-crud/models"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.Item{}, &models.User{})

	r := gin.Default()

	r.POST("/items", controllers.CreateItem)
	r.GET("/items", controllers.GetItems)
	r.GET("/items/:id", controllers.GetItemByID)
	r.PUT("/items/:id", controllers.UpdateItem)
	r.DELETE("/items/:id", controllers.DeleteItem)

	r.Run()
}
