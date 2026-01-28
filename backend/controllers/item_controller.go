package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeypc/go-crud/config"
	"github.com/jeypc/go-crud/models"
)

func GetItems(c *gin.Context) {
	var items []models.Item

	config.DB.Find(&items)

	c.JSON(http.StatusOK, gin.H{
		"data": items,
	})
}

func GetItemByID(c *gin.Context) {
	var item models.Item

	id := c.Param("id")

	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Item not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": item,
	})
}

func CreateItem(c *gin.Context) {
	var item models.Item

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := config.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Item created successfully",
		"data":    item,
	})
}

func UpdateItem(c *gin.Context) {
	var item models.Item
	id := c.Param("id")

	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Item Not Found",
		})
		return
	}

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	config.DB.Save(&item)

	c.JSON(http.StatusOK, gin.H{
		"message": "Item updated successfully",
		"data":    item,
	})
}

func DeleteItem(c *gin.Context) {
	var item models.Item
	id := c.Param("id")

	if err := config.DB.Delete(&item, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Item deleted successfully",
	})
}
