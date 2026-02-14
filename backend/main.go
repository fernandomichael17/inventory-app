package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jeypc/go-crud/config"
	"github.com/jeypc/go-crud/controllers"
	"github.com/jeypc/go-crud/middlewares"
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

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	authorized := r.Group("/")
	authorized.Use(middlewares.AuthMiddleware())
	{
		authorized.GET("/items", controllers.GetItems)
		authorized.GET("/items/:id", controllers.GetItemByID)

		adminRoutes := authorized.Group("/")
		adminRoutes.Use(middlewares.AdminOnly())
		{
			adminRoutes.POST("/items", controllers.CreateItem)
			adminRoutes.PUT("/items/:id", controllers.UpdateItem)
			adminRoutes.DELETE("/items/:id", controllers.DeleteItem)
		}

	}

	r.Run()
}
