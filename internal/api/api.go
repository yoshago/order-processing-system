package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yoshago/order-processing-system/internal/models"
	"gorm.io/gorm"
)

func InitRoutes(r *gin.Engine, database *gorm.DB) {
	r.POST("/orders", func(c *gin.Context) {
		var newOrder models.Order
		if err := c.ShouldBindJSON(&newOrder); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newOrder.ID = uuid.New().String()
		newOrder.Status = "pending"
		if err := database.Create(&newOrder).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
			return
		}
		c.JSON(http.StatusCreated, newOrder)
	})

	r.GET("/orders/:id", func(c *gin.Context) {
		var order models.Order
		if err := database.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}
		c.JSON(http.StatusOK, order)
	})

	// Endpoint to list all orders with pagination
	r.GET("/orders", func(c *gin.Context) {
		var orders []models.Order
		page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err != nil || page < 1 {
			page = 1
		}
		pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
		if err != nil || pageSize < 1 {
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		if err := database.Limit(pageSize).Offset(offset).Find(&orders).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
			return
		}

		c.JSON(http.StatusOK, orders)
	})
}
